package service

import (
	"context"

	"github.com/json-multiplex/iam/internal/models"
)

type Service interface {
	CreateAccount(context.Context, CreateAccountRequest) (models.Account, error)
	ListUsers(context.Context, ListUsersRequest) (ListUsersResponse, error)
	GetUser(context.Context, GetUserRequest) (models.User, error)
	CreateUser(context.Context, CreateUserRequest) (models.User, error)
	CreateSession(context.Context, CreateSessionRequest) (models.Session, error)
}

type CreateAccountRequest struct {
	Account      models.Account
	RootPassword string
}

type ListUsersRequest struct {
	Token string
}

type ListUsersResponse struct {
	Users []models.User
}

type GetUserRequest struct {
	ID    string
	Token string
}

type CreateUserRequest struct {
	User  models.User
	Token string
}

type CreateSessionRequest struct {
	Session models.Session
}
