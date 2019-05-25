package models

import "time"

type SAMLUser struct {
	IdentityProviderID string
	ID                 string
	CreateTime         time.Time
	UpdateTime         time.Time
}
