package models

import "time"

type AccessKey struct {
	UserID     string
	ID         string
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime *time.Time
	Secret     string
}
