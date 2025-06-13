package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id" `
	Status      string    `json:"status"`
	DueDate     string    `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	ListID      uuid.UUID `json:"list_id"`
	List        *List     `gorm:"foreignKey:ListID"`

	Tags []Tag `gorm:"many2many:task_tags;"`
}
