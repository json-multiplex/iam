package service

import (
	"context"

	"github.com/json-multiplex/iam/internal/models"
)

type Service interface {
	CreateAccount(context.Context, CreateAccountRequest) (models.Account, error)
	ListIdentityProviders(context.Context, ListIdentityProvidersRequest) (ListIdentityProvidersResponse, error)
	GetIdentityProvider(context.Context, GetIdentityProviderRequest) (models.IdentityProvider, error)
	CreateIdentityProvider(context.Context, CreateIdentityProviderRequest) (models.IdentityProvider, error)
	DeleteIdentityProvider(context.Context, DeleteIdentityProviderRequest) error
	ListUsers(context.Context, ListUsersRequest) (ListUsersResponse, error)
	GetUser(context.Context, GetUserRequest) (models.User, error)
	CreateUser(context.Context, CreateUserRequest) (models.User, error)
	DeleteUser(context.Context, DeleteUserRequest) error
	ListAccessKeys(context.Context, ListAccessKeysRequest) (ListAccessKeysResponse, error)
	GetAccessKey(context.Context, GetAccessKeyRequest) (models.AccessKey, error)
	CreateAccessKey(context.Context, CreateAccessKeyRequest) (models.AccessKey, error)
	DeleteAccessKey(context.Context, DeleteAccessKeyRequest) error
	CreateSession(context.Context, CreateSessionRequest) (models.Session, error)
}

type CreateAccountRequest struct {
	Account      models.Account
	RootPassword string
}

type ListIdentityProvidersRequest struct {
	Token string
}

type ListIdentityProvidersResponse struct {
	IdentityProviders []models.IdentityProvider
}

type GetIdentityProviderRequest struct {
	ID    string
	Token string
}

type CreateIdentityProviderRequest struct {
	IdentityProvider models.IdentityProvider
	Token            string
}

type DeleteIdentityProviderRequest struct {
	ID    string
	Token string
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

type DeleteUserRequest struct {
	ID    string
	Token string
}

type ListAccessKeysRequest struct {
	UserID string
	Token  string
}

type ListAccessKeysResponse struct {
	AccessKeys []models.AccessKey
}

type GetAccessKeyRequest struct {
	UserID string
	ID     string
	Token  string
}

type CreateAccessKeyRequest struct {
	AccessKey models.AccessKey
	Token     string
}

type DeleteAccessKeyRequest struct {
	UserID string
	ID     string
	Token  string
}

type CreateSessionRequest struct {
	Session models.Session
}
