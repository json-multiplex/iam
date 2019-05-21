package store

import (
	"context"
	"crypto/rand"
	"encoding/base64"
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

const AccessKeySecretsLength = 32

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

type dbIdentityProvider struct {
	ID              string     `db:"id"`
	CreateTime      time.Time  `db:"create_time"`
	UpdateTime      time.Time  `db:"update_time"`
	DeleteTime      *time.Time `db:"delete_time"`
	SAMLMetadataURL string     `db:"saml_metadata_url"`
	UserIDAttribute string     `db:"user_id_attribute"`
}

func (s *DBStore) ListIdentityProviders(ctx context.Context, in ListIdentityProvidersRequest) (ListIdentityProvidersResponse, error) {
	var dbIdentityProviders []dbIdentityProvider
	err := s.DB.SelectContext(ctx, &dbIdentityProviders, `
		SELECT
			id, create_time, update_time, delete_time, saml_metadata_url, user_id_attribute
		FROM
			identity_providers
		WHERE
			account_id = $1
	`, in.AccountID)

	if err != nil {
		return ListIdentityProvidersResponse{}, fmt.Errorf("failed to select identityProvider: %v", err)
	}

	identityProviders := make([]models.IdentityProvider, len(dbIdentityProviders))
	for i, identityProvider := range dbIdentityProviders {
		identityProviders[i] = models.IdentityProvider{
			ID:              identityProvider.ID,
			CreateTime:      identityProvider.CreateTime,
			UpdateTime:      identityProvider.UpdateTime,
			DeleteTime:      identityProvider.DeleteTime,
			SAMLMetadataURL: identityProvider.SAMLMetadataURL,
			UserIDAttribute: identityProvider.UserIDAttribute,
		}
	}

	return ListIdentityProvidersResponse{IdentityProviders: identityProviders}, nil
}

func (s *DBStore) GetIdentityProvider(ctx context.Context, in GetIdentityProviderRequest) (models.IdentityProvider, error) {
	var dbIdentityProvider dbIdentityProvider
	err := s.DB.GetContext(ctx, &dbIdentityProvider, `
		SELECT
			id, create_time, update_time, delete_time, saml_metadata_url, user_id_attribute
		FROM
			identity_providers
		WHERE
			account_id = $1 AND id = $2
	`, in.AccountID, in.ID)

	if err != nil {
		return models.IdentityProvider{}, fmt.Errorf("failed to select identityProvider: %v", err)
	}

	return models.IdentityProvider{
		ID:              dbIdentityProvider.ID,
		CreateTime:      dbIdentityProvider.CreateTime,
		UpdateTime:      dbIdentityProvider.UpdateTime,
		DeleteTime:      dbIdentityProvider.DeleteTime,
		SAMLMetadataURL: dbIdentityProvider.SAMLMetadataURL,
		UserIDAttribute: dbIdentityProvider.UserIDAttribute,
	}, nil
}

func (s *DBStore) CreateIdentityProvider(ctx context.Context, in CreateIdentityProviderRequest) (models.IdentityProvider, error) {
	createTime := time.Now()

	_, err := s.DB.ExecContext(ctx, `
		INSERT INTO identity_providers
			(account_id, id, create_time, update_time, delete_time, saml_metadata_url, user_id_attribute)
		VALUES
			($1, $2, $3, $3, NULL, $4, $5)
	`, in.AccountID, in.IdentityProvider.ID, createTime, in.IdentityProvider.SAMLMetadataURL, in.IdentityProvider.UserIDAttribute)

	if err != nil {
		return models.IdentityProvider{}, fmt.Errorf("failed to insert identityProvider into db: %v", err)
	}

	return models.IdentityProvider{
		ID:              in.IdentityProvider.ID,
		CreateTime:      createTime,
		UpdateTime:      createTime,
		SAMLMetadataURL: in.IdentityProvider.SAMLMetadataURL,
		UserIDAttribute: in.IdentityProvider.UserIDAttribute,
	}, nil
}

func (s *DBStore) DeleteIdentityProvider(ctx context.Context, in DeleteIdentityProviderRequest) error {
	_, err := s.DB.ExecContext(ctx, `
		DELETE FROM identity_providers WHERE account_id = $1 AND id = $2
	`, in.AccountID, in.ID)

	if err != nil {
		return fmt.Errorf("failed to delete identityProvider from db: %v", err)
	}

	return nil
}

type dbUser struct {
	ID           string     `db:"id"`
	CreateTime   time.Time  `db:"create_time"`
	UpdateTime   time.Time  `db:"update_time"`
	DeleteTime   *time.Time `db:"delete_time"`
	PasswordHash string     `db:"password_hash"`
}

func (s *DBStore) ListUsers(ctx context.Context, in ListUsersRequest) (ListUsersResponse, error) {
	var dbUsers []dbUser
	err := s.DB.SelectContext(ctx, &dbUsers, `
		SELECT
			id, create_time, update_time, delete_time
		FROM
			users
		WHERE
			account_id = $1
	`, in.AccountID)

	if err != nil {
		return ListUsersResponse{}, fmt.Errorf("failed to select users: %v", err)
	}

	users := make([]models.User, len(dbUsers))
	for i, user := range dbUsers {
		users[i] = models.User{
			ID:         user.ID,
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
			DeleteTime: user.DeleteTime,
		}
	}

	return ListUsersResponse{Users: users}, nil
}

func (s *DBStore) GetUser(ctx context.Context, in GetUserRequest) (models.User, error) {
	var dbUser dbUser
	err := s.DB.GetContext(ctx, &dbUser, `
	SELECT
		id, create_time, update_time, delete_time
	FROM
		users
	WHERE
		account_id = $1 AND id = $2
	`, in.AccountID, in.ID)

	if err != nil {
		return models.User{}, fmt.Errorf("failed to select user: %v", err)
	}

	return models.User{
		ID:         dbUser.ID,
		CreateTime: dbUser.CreateTime,
		UpdateTime: dbUser.UpdateTime,
		DeleteTime: dbUser.DeleteTime,
	}, nil
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

func (s *DBStore) DeleteUser(ctx context.Context, in DeleteUserRequest) error {
	_, err := s.DB.ExecContext(ctx, `
		DELETE FROM users WHERE account_id = $1 AND id = $2
	`, in.AccountID, in.ID)

	if err != nil {
		return fmt.Errorf("failed to delete user from db: %v", err)
	}

	return nil
}

type dbAccessKey struct {
	AccountID  uuid.UUID  `db:"account_id"`
	UserID     string     `db:"user_id"`
	ID         string     `db:"id"`
	CreateTime time.Time  `db:"create_time"`
	UpdateTime time.Time  `db:"update_time"`
	DeleteTime *time.Time `db:"delete_time"`
	SecretHash string     `db:"secret_hash"`
}

func (s *DBStore) ListAccessKeys(ctx context.Context, in ListAccessKeysRequest) (ListAccessKeysResponse, error) {
	var dbAccessKeys []dbAccessKey
	err := s.DB.SelectContext(ctx, &dbAccessKeys, `
		SELECT
			account_id, user_id, id, create_time, update_time, delete_time
		FROM
			access_keys
		WHERE
			account_id = $1 and user_id = $2
	`, in.AccountID, in.UserID)

	if err != nil {
		return ListAccessKeysResponse{}, fmt.Errorf("failed to select access keys: %v", err)
	}

	accessKeys := make([]models.AccessKey, len(dbAccessKeys))
	for i, accessKey := range dbAccessKeys {
		accessKeys[i] = models.AccessKey{
			UserID:     accessKey.UserID,
			ID:         accessKey.ID,
			CreateTime: accessKey.CreateTime,
			UpdateTime: accessKey.UpdateTime,
			DeleteTime: accessKey.DeleteTime,
		}
	}

	return ListAccessKeysResponse{AccessKeys: accessKeys}, nil
}

func (s *DBStore) GetAccessKey(ctx context.Context, in GetAccessKeyRequest) (models.AccessKey, error) {
	var dbAccessKey dbAccessKey
	err := s.DB.GetContext(ctx, &dbAccessKey, `
		SELECT
			user_id, id, create_time, update_time, delete_time
		FROM
			access_keys
		WHERE
			account_id = $1 AND user_id = $2 AND id = $3
	`, in.AccountID, in.UserID, in.ID)

	if err != nil {
		return models.AccessKey{}, fmt.Errorf("failed to select access key: %v", err)
	}

	return models.AccessKey{
		UserID:     dbAccessKey.UserID,
		ID:         dbAccessKey.ID,
		CreateTime: dbAccessKey.CreateTime,
		UpdateTime: dbAccessKey.UpdateTime,
		DeleteTime: dbAccessKey.DeleteTime,
	}, nil
}

func (s *DBStore) CreateAccessKey(ctx context.Context, in CreateAccessKeyRequest) (models.AccessKey, error) {
	createTime := time.Now()

	secret := make([]byte, AccessKeySecretsLength)
	_, err := rand.Read(secret)
	if err != nil {
		return models.AccessKey{}, fmt.Errorf("failed to read rand: %v", err)
	}

	secretBase64 := base64.StdEncoding.EncodeToString(secret)
	secretHash, err := bcrypt.GenerateFromPassword(secret, bcrypt.DefaultCost)
	if err != nil {
		return models.AccessKey{}, fmt.Errorf("failed to hash secret: %v", err)
	}

	_, err = s.DB.ExecContext(ctx, `
		INSERT INTO access_keys
			(account_id, user_id, id, create_time, update_time, delete_time, secret_hash)
		VALUES
			($1, $2, $3, $4, $4, NULL, $5);
	`, in.AccountID, in.AccessKey.UserID, in.AccessKey.ID, createTime, secretHash)

	if err != nil {
		return models.AccessKey{}, fmt.Errorf("failed to insert access key: %v", err)
	}

	return models.AccessKey{
		UserID:     in.AccessKey.UserID,
		ID:         in.AccessKey.ID,
		CreateTime: createTime,
		UpdateTime: createTime,
		Secret:     secretBase64,
	}, nil
}

func (s *DBStore) DeleteAccessKey(ctx context.Context, in DeleteAccessKeyRequest) error {
	_, err := s.DB.ExecContext(ctx, `
		DELETE FROM access_keys WHERE account_id = $1 AND user_id = $2 AND id = $3
	`, in.AccountID, in.UserID, in.ID)

	if err != nil {
		return fmt.Errorf("failed to delete access key: %v", err)
	}

	return nil
}

func (s *DBStore) CreateSession(ctx context.Context, in CreateSessionRequest) (models.Session, error) {
	if in.Session.AccessKeyID == "" {
		var dbUser dbUser
		err := s.DB.GetContext(ctx, &dbUser, `
			SELECT password_hash FROM users WHERE account_id = $1 AND id = $2;
		`, in.Session.AccountID, in.Session.UserID)

		if err != nil {
			return models.Session{}, fmt.Errorf("failed to get user: %v", err)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(in.Session.Secret)); err != nil {
			return models.Session{}, fmt.Errorf("error comparing hash and password: %v", err)
		}
	} else {
		var dbAccessKey dbAccessKey
		err := s.DB.GetContext(ctx, &dbAccessKey, `
			SELECT secret_hash FROM access_keys WHERE account_id = $1 AND user_id = $2 AND id = $3
		`, in.Session.AccountID, in.Session.UserID, in.Session.AccessKeyID)

		if err != nil {
			return models.Session{}, fmt.Errorf("failed to get access key: %v", err)
		}

		secret, err := base64.StdEncoding.DecodeString(in.Session.Secret)
		if err != nil {
			return models.Session{}, fmt.Errorf("failed to parse secret: %v", err)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbAccessKey.SecretHash), secret); err != nil {
			return models.Session{}, fmt.Errorf("error comparing hash and secret: %v", err)
		}
	}

	return models.Session{
		ID:        uuid.New().String(),
		AccountID: in.Session.AccountID,
		UserID:    in.Session.UserID,
	}, nil
}
