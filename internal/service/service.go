package service

import (
	"context"

	"github.com/json-multiplex/iam/internal/models"
)

type Service interface {
	CreateAccount(context.Context, CreateAccountRequest) (models.Account, error)
	CreateUser(context.Context, CreateUserRequest) (models.User, error)
	CreateSession(context.Context, CreateSessionRequest) (models.Session, error)
}

type CreateAccountRequest struct {
	Account      models.Account
	RootPassword string
}

type CreateUserRequest struct {
	User  models.User
	Token string
}

type CreateSessionRequest struct {
	Session models.Session
}
