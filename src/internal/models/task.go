package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	UserID      uuid.UUID  `json:"user_id"`
	Status      string     `json:"status"`
	IsCompleted bool       `json:"is_completed" gorm:"default:false"`
	DueDate     string     `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	ListID      uuid.UUID  `json:"list_id"`
	List        *List      `gorm:"foreignKey:ListID"`
	Tags        []Tag      `gorm:"many2many:task_tags;"`
	ParentID    *uuid.UUID `json:"parent_id" gorm:"index"`
	Subtasks    []Task     `gorm:"foreignKey:ParentID"`
	ColumnID    *uuid.UUID `json:"column_id" gorm:"index"`
	Position    int        `json:"position" gorm:"default:0"`
}
