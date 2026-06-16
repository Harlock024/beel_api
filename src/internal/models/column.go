package models

import (
	"time"

	"github.com/google/uuid"
)

type Column struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	BoardID   uuid.UUID `json:"board_id" gorm:"index;not null"`
	Board     Board     `gorm:"foreignKey:BoardID"`
	Title     string    `json:"title" gorm:"not null"`
	Position  int       `json:"position" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at"`
	Tasks     []Task    `gorm:"foreignKey:ColumnID"`
}
