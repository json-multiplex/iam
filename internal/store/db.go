package store

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/json-multiplex/iam/internal/models"
	"golang.org/x/crypto/bcrypt"
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
		return models.Account{}, fmt.Errorf("failed to insert account into db: %v", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.RootPassword), bcrypt.DefaultCost)
	if err != nil {
		return models.Account{}, fmt.Errorf("failed to bcrypt root password: %v", passwordHash)
	}

	_, err = s.DB.ExecContext(ctx, `
		INSERT INTO users
			(account_id, id, create_time, update_time, delete_time, is_root, password_hash)
		VALUES
			($1, 'root', $2, $2, NULL, TRUE, $3);
	`, id, createTime, passwordHash)

	if err != nil {
		return models.Account{}, fmt.Errorf("failed to insert root user into db: %v", err)
	}

	return models.Account{
		ID:         id,
		CreateTime: createTime,
		UpdateTime: createTime,
	}, nil
}
