package models

import (
	"time"
)

type Session struct {
	ID         string
	AccountID  string
	UserID     string
	CreateTime time.Time
	ExpireTime time.Time
	Password   string
	Token      string
}
