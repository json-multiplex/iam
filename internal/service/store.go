package service

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/json-multiplex/iam/internal/models"
	"github.com/json-multiplex/iam/internal/store"
)

type StoreService struct {
	Store                 store.Store
	TokenExpirationPeriod time.Duration
	TokenSignKey          *rsa.PrivateKey
	TokenVerifyKey        *rsa.PublicKey
}

func (s *StoreService) CreateAccount(ctx context.Context, in CreateAccountRequest) (models.Account, error) {
	return s.Store.CreateAccount(ctx, store.CreateAccountRequest{
		Account:      in.Account,
		RootPassword: in.RootPassword,
	})
}

func (s *StoreService) CreateSession(ctx context.Context, in CreateSessionRequest) (models.Session, error) {
	session, err := s.Store.CreateSession(ctx, store.CreateSessionRequest{Session: in.Session})
	if err != nil {
		return models.Session{}, fmt.Errorf("error from store: %v", err)
	}

	createTime := time.Now()
	expireTime := createTime.Add(s.TokenExpirationPeriod)

	session.Password = ""

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &jwt.StandardClaims{
		Subject:   in.Session.UserID,
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
