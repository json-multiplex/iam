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

	srv := server{
		Service: &service.StoreService{
			Store:                 &store.DBStore{DB: dbConn},
			TokenExpirationPeriod: 24 * time.Hour,
			TokenSignKey:          tokenSignKey,
			TokenVerifyKey:        tokenVerifyKey,
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

func (s *server) CreateSession(ctx context.Context, in *pb.CreateSessionRequest) (*pb.Session, error) {
	accountID := strings.Split(in.Session.Account, "/")[1]
	userID := strings.Split(in.Session.User, "/")[1]

	session, err := s.Service.CreateSession(ctx, service.CreateSessionRequest{Session: models.Session{
		AccountID: accountID,
		UserID:    userID,
		Password:  in.Session.Password,
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

	return &pb.Session{
		Name:       fmt.Sprintf("sessions/%s", session.ID),
		Account:    fmt.Sprintf("accounts/%s", session.AccountID),
		User:       fmt.Sprintf("users/%s", session.UserID),
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
