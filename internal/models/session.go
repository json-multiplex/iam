package models

import (
	"time"
)

type Session struct {
	ID          string
	AccountID   string
	UserID      string
	AccessKeyID string
	CreateTime  time.Time
	ExpireTime  time.Time
	Secret      string
	Token       string
}
