package models

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id"`
	Status      bool      `json:"status"`
}
