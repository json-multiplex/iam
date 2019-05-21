package service

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/json-multiplex/iam/internal/models"
	"github.com/json-multiplex/iam/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StoreService struct {
	Store                 store.Store
	TokenExpirationPeriod time.Duration
	TokenSignKey          *rsa.PrivateKey
	TokenVerifyKey        *rsa.PublicKey
}

func (s *StoreService) CreateAccount(ctx context.Context, in CreateAccountRequest) (models.Account, error) {
	account, err := s.Store.CreateAccount(ctx, store.CreateAccountRequest{
		Account:      in.Account,
		RootPassword: in.RootPassword,
	})

	if err != nil {
		return models.Account{}, fmt.Errorf("error from store: %v", err)
	}

	return account, err
}

func (s *StoreService) ListIdentityProviders(ctx context.Context, in ListIdentityProvidersRequest) (ListIdentityProvidersResponse, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return ListIdentityProvidersResponse{}, err
	}

	identityProvidersList, err := s.Store.ListIdentityProviders(ctx, store.ListIdentityProvidersRequest{AccountID: claims.Audience})
	if err != nil {
		return ListIdentityProvidersResponse{}, fmt.Errorf("error from store: %v", err)
	}

	return ListIdentityProvidersResponse{IdentityProviders: identityProvidersList.IdentityProviders}, nil
}

func (s *StoreService) GetIdentityProvider(ctx context.Context, in GetIdentityProviderRequest) (models.IdentityProvider, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return models.IdentityProvider{}, err
	}

	identityProvider, err := s.Store.GetIdentityProvider(ctx, store.GetIdentityProviderRequest{
		AccountID: claims.Audience,
		ID:        in.ID,
	})

	if err != nil {
		return models.IdentityProvider{}, fmt.Errorf("error from store: %v", err)
	}

	return identityProvider, nil
}

func (s *StoreService) CreateIdentityProvider(ctx context.Context, in CreateIdentityProviderRequest) (models.IdentityProvider, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return models.IdentityProvider{}, err
	}

	return s.Store.CreateIdentityProvider(ctx, store.CreateIdentityProviderRequest{
		AccountID:        claims.Audience,
		IdentityProvider: in.IdentityProvider,
	})
}

func (s *StoreService) DeleteIdentityProvider(ctx context.Context, in DeleteIdentityProviderRequest) error {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return err
	}

	err = s.Store.DeleteIdentityProvider(ctx, store.DeleteIdentityProviderRequest{
		AccountID: claims.Audience,
		ID:        in.ID,
	})

	if err != nil {
		return fmt.Errorf("error from store: %v", err)
	}

	return nil
}

func (s *StoreService) ListUsers(ctx context.Context, in ListUsersRequest) (ListUsersResponse, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return ListUsersResponse{}, err
	}

	usersList, err := s.Store.ListUsers(ctx, store.ListUsersRequest{AccountID: claims.Audience})
	if err != nil {
		return ListUsersResponse{}, fmt.Errorf("error from store: %v", err)
	}

	return ListUsersResponse{Users: usersList.Users}, nil
}

func (s *StoreService) GetUser(ctx context.Context, in GetUserRequest) (models.User, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.Store.GetUser(ctx, store.GetUserRequest{
		AccountID: claims.Audience,
		ID:        in.ID,
	})

	if err != nil {
		return models.User{}, fmt.Errorf("error from store: %v", err)
	}

	return user, nil
}

func (s *StoreService) CreateUser(ctx context.Context, in CreateUserRequest) (models.User, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return models.User{}, err
	}

	return s.Store.CreateUser(ctx, store.CreateUserRequest{
		AccountID: claims.Audience,
		User:      in.User,
	})
}

func (s *StoreService) DeleteUser(ctx context.Context, in DeleteUserRequest) error {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return err
	}

	err = s.Store.DeleteUser(ctx, store.DeleteUserRequest{
		AccountID: claims.Audience,
		ID:        in.ID,
	})

	if err != nil {
		return fmt.Errorf("error from store: %v", err)
	}

	return nil
}

func (s *StoreService) ListAccessKeys(ctx context.Context, in ListAccessKeysRequest) (ListAccessKeysResponse, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return ListAccessKeysResponse{}, err
	}

	accessKeysList, err := s.Store.ListAccessKeys(ctx, store.ListAccessKeysRequest{
		AccountID: claims.Audience,
		UserID:    in.UserID,
	})

	if err != nil {
		return ListAccessKeysResponse{}, fmt.Errorf("error from store: %v", err)
	}

	return ListAccessKeysResponse{AccessKeys: accessKeysList.AccessKeys}, nil
}

func (s *StoreService) GetAccessKey(ctx context.Context, in GetAccessKeyRequest) (models.AccessKey, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return models.AccessKey{}, err
	}

	accessKey, err := s.Store.GetAccessKey(ctx, store.GetAccessKeyRequest{
		AccountID: claims.Audience,
		UserID:    in.UserID,
		ID:        in.ID,
	})

	if err != nil {
		return models.AccessKey{}, fmt.Errorf("error from store: %v", err)
	}

	return accessKey, nil
}

func (s *StoreService) CreateAccessKey(ctx context.Context, in CreateAccessKeyRequest) (models.AccessKey, error) {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return models.AccessKey{}, err
	}

	return s.Store.CreateAccessKey(ctx, store.CreateAccessKeyRequest{
		AccountID: claims.Audience,
		AccessKey: in.AccessKey,
	})
}

func (s *StoreService) DeleteAccessKey(ctx context.Context, in DeleteAccessKeyRequest) error {
	claims, err := s.parseToken(in.Token)
	if err != nil {
		return err
	}

	err = s.Store.DeleteAccessKey(ctx, store.DeleteAccessKeyRequest{
		AccountID: claims.Audience,
		UserID:    in.UserID,
		ID:        in.ID,
	})

	if err != nil {
		return fmt.Errorf("error from store: %v", err)
	}

	return nil
}

func (s *StoreService) CreateSession(ctx context.Context, in CreateSessionRequest) (models.Session, error) {
	session, err := s.Store.CreateSession(ctx, store.CreateSessionRequest{Session: in.Session})
	if err != nil {
		return models.Session{}, fmt.Errorf("error from store: %v", err)
	}

	createTime := time.Now()
	expireTime := createTime.Add(s.TokenExpirationPeriod)

	var subject string
	if in.Session.AccessKeyID == "" {
		subject = fmt.Sprintf("users/%s", in.Session.UserID)
	} else {
		subject = fmt.Sprintf("users/%s/accessKeys/%s", in.Session.UserID, in.Session.AccessKeyID)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &jwt.StandardClaims{
		Subject:   subject,
		Audience:  in.Session.AccountID,
		IssuedAt:  createTime.Unix(),
		ExpiresAt: expireTime.Unix(),
	})

	tokenString, err := token.SignedString(s.TokenSignKey)
	if err != nil {
		return models.Session{}, fmt.Errorf("error signing token: %v", err)
	}

	return models.Session{
		ID:         session.ID,
		AccountID:  session.AccountID,
		UserID:     session.UserID,
		CreateTime: createTime,
		ExpireTime: expireTime,
		Token:      tokenString,
	}, nil
}

func (s *StoreService) parseToken(token string) (*jwt.StandardClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			msg := fmt.Sprintf("unexpected token signing method: %v", token.Header["alg"])
			return nil, status.Error(codes.Unauthenticated, msg)
		}

		return s.TokenVerifyKey, nil
	})

	return parsed.Claims.(*jwt.StandardClaims), err
}
