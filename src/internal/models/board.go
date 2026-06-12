package models

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"index;not null"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Columns   []Column  `gorm:"foreignKey:BoardID"`
}
