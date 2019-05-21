package models

import "time"

type IdentityProvider struct {
	ID              string
	CreateTime      time.Time
	UpdateTime      time.Time
	DeleteTime      *time.Time
	SAMLMetadataURL string
	UserIDAttribute string
}
