package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jmoiron/sqlx"
	pb "github.com/json-multiplex/iam/generated/jsonmultiplex/iam/v0"
	"github.com/json-multiplex/iam/internal/models"
	"github.com/json-multiplex/iam/internal/service"
	"github.com/json-multiplex/iam/internal/store"
	_ "github.com/lib/pq"
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "IAM_SERVICE", 0)
	db := fs.String("db", "", "database url")
	tokenSignKeyPEM := fs.String("token_sign_key", "", "PEM-encoded key for signing tokens")
	tokenVerifyKeyPEM := fs.String("token_verify_key", "", "PEM-encoded key for verifying tokens")
	samlVerifyKeyPEM := fs.String("saml_verify_key", "", "PEM-encoded key for verifying SAML tokens")
	fs.Parse(os.Args[1:])

	dbConn, err := sqlx.Open("postgres", *db)
	if err != nil {
		log.Fatalf("could not open db connection: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	var tokenSignKey *rsa.PrivateKey
	var tokenVerifyKey *rsa.PublicKey
	var samlVerifyKey *rsa.PublicKey

	if block, _ := pem.Decode([]byte(*tokenSignKeyPEM)); block != nil {
		tokenSignKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatalf("error parsing token sign key: %v", err)
		}
	}

	if block, _ := pem.Decode([]byte(*tokenVerifyKeyPEM)); block != nil {
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			log.Fatalf("error parsing token verify key: %v", err)
		}

		var ok bool
		tokenVerifyKey, ok = key.(*rsa.PublicKey)
		if !ok {
			log.Fatalf("verify key must be RSA: %v", err)
		}
	}

	if block, _ := pem.Decode([]byte(*samlVerifyKeyPEM)); block != nil {
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			log.Fatalf("error parsing saml verify key: %v", err)
		}

		var ok bool
		samlVerifyKey, ok = key.(*rsa.PublicKey)
		if !ok {
			log.Fatalf("verify key must be RSA: %v", err)
		}
	}

	srv := server{
		Service: &service.StoreService{
			Store:                 &store.DBStore{DB: dbConn},
			TokenExpirationPeriod: 24 * time.Hour,
			TokenSignKey:          tokenSignKey,
			TokenVerifyKey:        tokenVerifyKey,
			SAMLVerifyKey:         samlVerifyKey,
		},
	}

	grpcServer := grpc.NewServer()
	pb.RegisterIAMServer(grpcServer, &srv)
	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("could not listen for tcp: %v", err)
	}

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("error serving: %v", err)
	}
}

type server struct {
	Service service.Service
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	account, err := s.Service.CreateAccount(ctx, service.CreateAccountRequest{
		RootPassword: in.Account.RootPassword,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(account.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(account.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if !account.DeleteTime.IsZero() {
		var err error
		deleteTime, err = ptypes.TimestampProto(account.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Account{
		Name:       fmt.Sprintf("accounts/%s", account.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
		DeleteTime: deleteTime,
	}, nil
}

func (s *server) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	usersList, err := s.Service.ListUsers(ctx, service.ListUsersRequest{Token: token})
	if err != nil {
		return nil, err
	}

	users := make([]*pb.User, len(usersList.Users))
	for i, user := range usersList.Users {
		createTime, err := ptypes.TimestampProto(user.CreateTime)
		if err != nil {
			return nil, err
		}

		updateTime, err := ptypes.TimestampProto(user.UpdateTime)
		if err != nil {
			return nil, err
		}

		var deleteTime *timestamp.Timestamp
		if user.DeleteTime != nil {
			var err error
			deleteTime, err = ptypes.TimestampProto(*user.DeleteTime)
			if err != nil {
				return nil, err
			}
		}

		users[i] = &pb.User{
			Name:       fmt.Sprintf("users/%s", user.ID),
			CreateTime: createTime,
			UpdateTime: updateTime,
			DeleteTime: deleteTime,
		}
	}

	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.Split(in.Name, "/")[1]
	user, err := s.Service.GetUser(ctx, service.GetUserRequest{
		ID:    userID,
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(user.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(user.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if user.DeleteTime != nil {
		var err error
		deleteTime, err = ptypes.TimestampProto(*user.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.User{
		Name:       fmt.Sprintf("users/%s", user.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
		DeleteTime: deleteTime,
	}, nil
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.Split(in.User.Name, "/")[1]
	user, err := s.Service.CreateUser(ctx, service.CreateUserRequest{
		User: models.User{
			ID:       userID,
			Password: in.User.Password,
		},
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(user.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(user.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if user.DeleteTime != nil {
		var err error
		deleteTime, err = ptypes.TimestampProto(*user.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.User{
		Name:       fmt.Sprintf("users/%s", user.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
		DeleteTime: deleteTime,
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*empty.Empty, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.Split(in.Name, "/")[1]

	err = s.Service.DeleteUser(ctx, service.DeleteUserRequest{
		Token: token,
		ID:    userID,
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *server) ListAccessKeys(ctx context.Context, in *pb.ListAccessKeysRequest) (*pb.ListAccessKeysResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.Split(in.Parent, "/")[1]
	accessKeysList, err := s.Service.ListAccessKeys(ctx, service.ListAccessKeysRequest{
		UserID: userID,
		Token:  token,
	})

	if err != nil {
		return nil, err
	}

	accessKeys := make([]*pb.AccessKey, len(accessKeysList.AccessKeys))
	for i, accessKey := range accessKeysList.AccessKeys {
		createTime, err := ptypes.TimestampProto(accessKey.CreateTime)
		if err != nil {
			return nil, err
		}

		updateTime, err := ptypes.TimestampProto(accessKey.UpdateTime)
		if err != nil {
			return nil, err
		}

		var deleteTime *timestamp.Timestamp
		if accessKey.DeleteTime != nil {
			var err error
			deleteTime, err = ptypes.TimestampProto(*accessKey.DeleteTime)
			if err != nil {
				return nil, err
			}
		}

		accessKeys[i] = &pb.AccessKey{
			Name:       fmt.Sprintf("users/%s/accessKeys/%s", accessKey.UserID, accessKey.ID),
			CreateTime: createTime,
			UpdateTime: updateTime,
			DeleteTime: deleteTime,
		}
	}

	return &pb.ListAccessKeysResponse{AccessKeys: accessKeys}, nil
}

func (s *server) GetAccessKey(ctx context.Context, in *pb.GetAccessKeyRequest) (*pb.AccessKey, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.Split(in.Name, "/")[1]
	accessKeyID := strings.Split(in.Name, "/")[3]
	accessKey, err := s.Service.GetAccessKey(ctx, service.GetAccessKeyRequest{
		UserID: userID,
		ID:     accessKeyID,
		Token:  token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(accessKey.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(accessKey.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if accessKey.DeleteTime != nil {
		var err error
		deleteTime, err = ptypes.TimestampProto(*accessKey.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.AccessKey{
		Name:       fmt.Sprintf("users/%s/accessKeys/%s", accessKey.UserID, accessKey.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
		DeleteTime: deleteTime,
	}, nil
}

func (s *server) CreateAccessKey(ctx context.Context, in *pb.CreateAccessKeyRequest) (*pb.AccessKey, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.Split(in.Parent, "/")[1]
	accessKey, err := s.Service.CreateAccessKey(ctx, service.CreateAccessKeyRequest{
		AccessKey: models.AccessKey{
			UserID: userID,
			ID:     in.AccessKeyId,
		},
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(accessKey.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(accessKey.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if accessKey.DeleteTime != nil {
		var err error
		deleteTime, err = ptypes.TimestampProto(*accessKey.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.AccessKey{
		Name:       fmt.Sprintf("users/%s/accessKeys/%s", accessKey.UserID, accessKey.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
		DeleteTime: deleteTime,
		Secret:     accessKey.Secret,
	}, nil
}

func (s *server) DeleteAccessKey(ctx context.Context, in *pb.DeleteAccessKeyRequest) (*empty.Empty, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	userID := strings.Split(in.Name, "/")[1]
	accessKeyID := strings.Split(in.Name, "/")[3]
	err = s.Service.DeleteAccessKey(ctx, service.DeleteAccessKeyRequest{
		UserID: userID,
		ID:     accessKeyID,
		Token:  token,
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *server) ListIdentityProviders(ctx context.Context, in *pb.ListIdentityProvidersRequest) (*pb.ListIdentityProvidersResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProvidersList, err := s.Service.ListIdentityProviders(ctx, service.ListIdentityProvidersRequest{Token: token})
	if err != nil {
		return nil, err
	}

	identityProviders := make([]*pb.IdentityProvider, len(identityProvidersList.IdentityProviders))
	for i, identityProvider := range identityProvidersList.IdentityProviders {
		createTime, err := ptypes.TimestampProto(identityProvider.CreateTime)
		if err != nil {
			return nil, err
		}

		updateTime, err := ptypes.TimestampProto(identityProvider.UpdateTime)
		if err != nil {
			return nil, err
		}

		var deleteTime *timestamp.Timestamp
		if identityProvider.DeleteTime != nil {
			var err error
			deleteTime, err = ptypes.TimestampProto(*identityProvider.DeleteTime)
			if err != nil {
				return nil, err
			}
		}

		identityProviders[i] = &pb.IdentityProvider{
			Name:            fmt.Sprintf("identityProviders/%s", identityProvider.ID),
			CreateTime:      createTime,
			UpdateTime:      updateTime,
			DeleteTime:      deleteTime,
			SamlMetadataUrl: identityProvider.SAMLMetadataURL,
			UserIdAttribute: identityProvider.UserIDAttribute,
		}
	}

	return &pb.ListIdentityProvidersResponse{IdentityProviders: identityProviders}, nil
}

func (s *server) GetIdentityProvider(ctx context.Context, in *pb.GetIdentityProviderRequest) (*pb.IdentityProvider, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProviderID := strings.Split(in.Name, "/")[1]
	identityProvider, err := s.Service.GetIdentityProvider(ctx, service.GetIdentityProviderRequest{
		ID:    identityProviderID,
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(identityProvider.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(identityProvider.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if identityProvider.DeleteTime != nil {
		var err error
		deleteTime, err = ptypes.TimestampProto(*identityProvider.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.IdentityProvider{
		Name:            fmt.Sprintf("identityProviders/%s", identityProvider.ID),
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DeleteTime:      deleteTime,
		SamlMetadataUrl: identityProvider.SAMLMetadataURL,
		UserIdAttribute: identityProvider.UserIDAttribute,
	}, nil
}

func (s *server) CreateIdentityProvider(ctx context.Context, in *pb.CreateIdentityProviderRequest) (*pb.IdentityProvider, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProviderID := strings.Split(in.IdentityProvider.Name, "/")[1]
	identityProvider, err := s.Service.CreateIdentityProvider(ctx, service.CreateIdentityProviderRequest{
		IdentityProvider: models.IdentityProvider{
			ID:              identityProviderID,
			SAMLMetadataURL: in.IdentityProvider.SamlMetadataUrl,
			UserIDAttribute: in.IdentityProvider.UserIdAttribute,
		},
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(identityProvider.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(identityProvider.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if identityProvider.DeleteTime != nil {
		var err error
		deleteTime, err = ptypes.TimestampProto(*identityProvider.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.IdentityProvider{
		Name:            fmt.Sprintf("identityProviders/%s", identityProvider.ID),
		CreateTime:      createTime,
		UpdateTime:      updateTime,
		DeleteTime:      deleteTime,
		SamlMetadataUrl: identityProvider.SAMLMetadataURL,
		UserIdAttribute: identityProvider.UserIDAttribute,
	}, nil
}

func (s *server) DeleteIdentityProvider(ctx context.Context, in *pb.DeleteIdentityProviderRequest) (*empty.Empty, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProviderID := strings.Split(in.Name, "/")[1]

	err = s.Service.DeleteIdentityProvider(ctx, service.DeleteIdentityProviderRequest{
		Token: token,
		ID:    identityProviderID,
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *server) ListSamlUsers(ctx context.Context, in *pb.ListSamlUsersRequest) (*pb.ListSamlUsersResponse, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProviderID := strings.Split(in.Parent, "/")[1]
	samlUsersList, err := s.Service.ListSAMLUsers(ctx, service.ListSAMLUsersRequest{
		IdentityProviderID: identityProviderID,
		Token:              token,
	})

	if err != nil {
		return nil, err
	}

	samlUsers := make([]*pb.SamlUser, len(samlUsersList.SAMLUsers))
	for i, samlUser := range samlUsersList.SAMLUsers {
		createTime, err := ptypes.TimestampProto(samlUser.CreateTime)
		if err != nil {
			return nil, err
		}

		updateTime, err := ptypes.TimestampProto(samlUser.UpdateTime)
		if err != nil {
			return nil, err
		}

		samlUsers[i] = &pb.SamlUser{
			Name:       fmt.Sprintf("identityProviders/%s/samlUsers/%s", samlUser.IdentityProviderID, samlUser.ID),
			CreateTime: createTime,
			UpdateTime: updateTime,
		}
	}

	return &pb.ListSamlUsersResponse{SamlUsers: samlUsers}, nil
}

func (s *server) GetSamlUser(ctx context.Context, in *pb.GetSamlUserRequest) (*pb.SamlUser, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProviderID := strings.Split(in.Name, "/")[1]
	samlUserID := strings.Split(in.Name, "/")[3]
	samlUser, err := s.Service.GetSAMLUser(ctx, service.GetSAMLUserRequest{
		IdentityProviderID: identityProviderID,
		ID:                 samlUserID,
		Token:              token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(samlUser.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(samlUser.UpdateTime)
	if err != nil {
		return nil, err
	}

	return &pb.SamlUser{
		Name:       fmt.Sprintf("identityProviders/%s/samlUsers/%s", samlUser.IdentityProviderID, samlUser.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
	}, nil
}

func (s *server) CreateSamlUser(ctx context.Context, in *pb.CreateSamlUserRequest) (*pb.SamlUser, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProviderID := strings.Split(in.Parent, "/")[1]
	samlUser, err := s.Service.CreateSAMLUser(ctx, service.CreateSAMLUserRequest{
		SAMLUser: models.SAMLUser{
			IdentityProviderID: identityProviderID,
			ID:                 in.SamlUserId,
		},
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(samlUser.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(samlUser.UpdateTime)
	if err != nil {
		return nil, err
	}

	return &pb.SamlUser{
		Name:       fmt.Sprintf("identityProviders/%s/samlUsers/%s", samlUser.IdentityProviderID, samlUser.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
	}, nil
}

func (s *server) DeleteSamlUser(ctx context.Context, in *pb.DeleteSamlUserRequest) (*empty.Empty, error) {
	token, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}

	identityProviderID := strings.Split(in.Name, "/")[1]
	samlUserID := strings.Split(in.Name, "/")[3]
	err = s.Service.DeleteSAMLUser(ctx, service.DeleteSAMLUserRequest{
		IdentityProviderID: identityProviderID,
		ID:                 samlUserID,
		Token:              token,
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *server) CreateSession(ctx context.Context, in *pb.CreateSessionRequest) (*pb.Session, error) {
	accountID := strings.Split(in.Session.Account, "/")[1]

	subjectSegments := strings.Split(in.Session.Subject, "/")
	userID := subjectSegments[1]
	var accessKeyID string
	if len(subjectSegments) == 4 {
		accessKeyID = subjectSegments[3]
	}

	session, err := s.Service.CreateSession(ctx, service.CreateSessionRequest{Session: models.Session{
		AccountID:   accountID,
		UserID:      userID,
		AccessKeyID: accessKeyID,
		Secret:      in.Session.Secret,
	}})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(session.CreateTime)
	if err != nil {
		return nil, err
	}

	expireTime, err := ptypes.TimestampProto(session.ExpireTime)
	if err != nil {
		return nil, err
	}

	var subject string
	if session.AccessKeyID == "" {
		subject = fmt.Sprintf("users/%s", session.UserID)
	} else {
		subject = fmt.Sprintf("users/%s/accessKeys/%s", session.UserID, session.AccessKeyID)
	}

	return &pb.Session{
		Name:       fmt.Sprintf("sessions/%s", session.ID),
		Account:    fmt.Sprintf("accounts/%s", session.AccountID),
		Subject:    subject,
		CreateTime: createTime,
		ExpireTime: expireTime,
		Token:      session.Token,
	}, nil
}

func (s *server) getToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no metadata provided")
	}

	authorization := md.Get("Authorization")
	if len(authorization) != 1 {
		return "", status.Error(codes.Unauthenticated, "no authorization provided")
	}

	authorizationSegments := strings.Split(authorization[0], " ")
	if len(authorizationSegments) != 2 || authorizationSegments[0] != "Bearer" {
		return "", status.Error(codes.Unauthenticated, "invalid authorization header")
	}

	return authorizationSegments[1], nil
}
