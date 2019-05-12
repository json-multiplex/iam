package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime time.Time
}
