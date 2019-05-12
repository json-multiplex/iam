package models

import "time"

type User struct {
	ID         string
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime time.Time
	Password   string
}
