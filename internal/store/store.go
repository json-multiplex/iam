package store

import (
	"context"

	"github.com/json-multiplex/iam/internal/models"
)

type Store interface {
	CreateAccount(context.Context, CreateAccountRequest) (models.Account, error)
	ListUsers(context.Context, ListUsersRequest) (ListUsersResponse, error)
	CreateUser(context.Context, CreateUserRequest) (models.User, error)
	CreateSession(context.Context, CreateSessionRequest) (models.Session, error)
}

type CreateAccountRequest struct {
	Account      models.Account
	RootPassword string
}

type ListUsersRequest struct {
	AccountID string
}

type ListUsersResponse struct {
	Users []models.User
}

type CreateUserRequest struct {
	AccountID string
	User      models.User
}

type CreateSessionRequest struct {
	Session models.Session
}
