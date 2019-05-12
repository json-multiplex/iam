package service

import (
	"context"

	"github.com/json-multiplex/iam/internal/models"
	"github.com/json-multiplex/iam/internal/store"
)

type StoreService struct {
	Store store.Store
}

func (s *StoreService) CreateAccount(ctx context.Context, in CreateAccountRequest) (models.Account, error) {
	return s.Store.CreateAccount(ctx, store.CreateAccountRequest{
		Account:      in.Account,
		RootPassword: in.RootPassword,
	})
}
