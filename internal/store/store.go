package store

import (
	"context"

	"github.com/json-multiplex/iam/internal/models"
)

type Store interface {
	CreateAccount(context.Context, CreateAccountRequest) (models.Account, error)
	ListUsers(context.Context, ListUsersRequest) (ListUsersResponse, error)
	GetUser(context.Context, GetUserRequest) (models.User, error)
	CreateUser(context.Context, CreateUserRequest) (models.User, error)
	DeleteUser(context.Context, DeleteUserRequest) error
	ListAccessKeys(context.Context, ListAccessKeysRequest) (ListAccessKeysResponse, error)
	GetAccessKey(context.Context, GetAccessKeyRequest) (models.AccessKey, error)
	CreateAccessKey(context.Context, CreateAccessKeyRequest) (models.AccessKey, error)
	DeleteAccessKey(context.Context, DeleteAccessKeyRequest) error
	ListIdentityProviders(context.Context, ListIdentityProvidersRequest) (ListIdentityProvidersResponse, error)
	GetIdentityProvider(context.Context, GetIdentityProviderRequest) (models.IdentityProvider, error)
	CreateIdentityProvider(context.Context, CreateIdentityProviderRequest) (models.IdentityProvider, error)
	DeleteIdentityProvider(context.Context, DeleteIdentityProviderRequest) error
	ListSAMLUsers(context.Context, ListSAMLUsersRequest) (ListSAMLUsersResponse, error)
	GetSAMLUser(context.Context, GetSAMLUserRequest) (models.SAMLUser, error)
	CreateSAMLUser(context.Context, CreateSAMLUserRequest) (models.SAMLUser, error)
	DeleteSAMLUser(context.Context, DeleteSAMLUserRequest) error
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

type GetUserRequest struct {
	AccountID string
	ID        string
}

type CreateUserRequest struct {
	AccountID string
	User      models.User
}

type DeleteUserRequest struct {
	AccountID string
	ID        string
}

type ListAccessKeysRequest struct {
	AccountID string
	UserID    string
}

type ListAccessKeysResponse struct {
	AccessKeys []models.AccessKey
}

type GetAccessKeyRequest struct {
	AccountID string
	UserID    string
	ID        string
}

type CreateAccessKeyRequest struct {
	AccountID string
	AccessKey models.AccessKey
}

type DeleteAccessKeyRequest struct {
	AccountID string
	UserID    string
	ID        string
}

type ListIdentityProvidersRequest struct {
	AccountID string
}

type ListIdentityProvidersResponse struct {
	IdentityProviders []models.IdentityProvider
}

type GetIdentityProviderRequest struct {
	AccountID string
	ID        string
}

type CreateIdentityProviderRequest struct {
	AccountID        string
	IdentityProvider models.IdentityProvider
}

type DeleteIdentityProviderRequest struct {
	AccountID string
	ID        string
}

type ListSAMLUsersRequest struct {
	AccountID          string
	IdentityProviderID string
}

type ListSAMLUsersResponse struct {
	SAMLUsers []models.SAMLUser
}

type GetSAMLUserRequest struct {
	AccountID          string
	IdentityProviderID string
	ID                 string
}

type CreateSAMLUserRequest struct {
	AccountID string
	SAMLUser  models.SAMLUser
}

type DeleteSAMLUserRequest struct {
	AccountID          string
	IdentityProviderID string
	ID                 string
}

type CreateSessionRequest struct {
	Session models.Session
}
