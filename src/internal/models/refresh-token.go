package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID          uuid.UUID `gorm:"primary_key" json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	HashedToken string    `json:"hashed_token"`
	IsRevoked   bool      `json:"is_revoked"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	User        User      `gorm:"foreignKey:UserID"`
}
