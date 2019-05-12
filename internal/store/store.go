package store

import (
	"context"

	"github.com/json-multiplex/iam/internal/models"
)

type Store interface {
	CreateAccount(context.Context, CreateAccountRequest) (models.Account, error)
}

type CreateAccountRequest struct {
	Account      models.Account
	RootPassword string
}
