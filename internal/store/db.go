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
		ID:         id.String(),
		CreateTime: createTime,
		UpdateTime: createTime,
	}, nil
}

type dbUser struct {
	PasswordHash string `db:"password_hash"`
}

func (s *DBStore) CreateUser(ctx context.Context, in CreateUserRequest) (models.User, error) {
	createTime := time.Now()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to bcrypt root password: %v", passwordHash)
	}

	_, err = s.DB.ExecContext(ctx, `
		INSERT INTO users
			(account_id, id, create_time, update_time, delete_time, is_root, password_hash)
		VALUES
			($1, $2, $3, $3, NULL, FALSE, $4);
	`, in.AccountID, in.User.ID, createTime, passwordHash)

	if err != nil {
		return models.User{}, fmt.Errorf("failed to insert user into db: %v", err)
	}

	return models.User{
		ID:         in.User.ID,
		CreateTime: createTime,
		UpdateTime: createTime,
	}, nil
}

func (s *DBStore) CreateSession(ctx context.Context, in CreateSessionRequest) (models.Session, error) {
	var user dbUser
	err := s.DB.GetContext(ctx, &user, `
		SELECT password_hash FROM users WHERE account_id = $1 AND id = $2;
	`, in.Session.AccountID, in.Session.UserID)

	if err != nil {
		return models.Session{}, fmt.Errorf("failed to get user: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Session.Password)); err != nil {
		return models.Session{}, fmt.Errorf("error comparing hash and password: %v", err)
	}

	return models.Session{
		ID:        uuid.New().String(),
		AccountID: in.Session.AccountID,
		UserID:    in.Session.UserID,
	}, err
}
