package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id"`
	Status      bool      `json:"status"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	ListID      *uuid.UUID
	List        *List `gorm:"foreignKey:ListID"`

	Tags []Tag `gorm:"many2many:task_tags;"`
}
