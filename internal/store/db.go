package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/json-multiplex/iam/internal/models"
	"github.com/pkg/errors"
)

type DBStore struct {
	DB *sqlx.DB
}

func (s *DBStore) CreateAccount(ctx context.Context, in CreateAccountRequest) (models.Account, error) {
	id := uuid.New()
	createTime := time.Now()

	_, err := s.DB.ExecContext(ctx, `
		INSERT INTO accounts
			(id, create_time, update_time, delete_time)
		VALUES
			($1, $2, $2, NULL);
	`, id, createTime)

	if err != nil {
		return models.Account{}, errors.Wrap(err, "failed to insert into db")
	}

	return models.Account{
		ID:         id,
		CreateTime: createTime,
		UpdateTime: createTime,
	}, nil
}
