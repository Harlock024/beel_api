package dtos

import (
	"github.com/google/uuid"
)

type NewTaskDTO struct {
	Title string `json:"title"`
}
type UpdateTaskDTO struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	IsCompleted bool        `json:"is_completed"`
	DueDate     string      `json:"due_date"`
	ListID      *uuid.UUID  `json:"list_id"`
	TagIDs      []uuid.UUID `json:"tag_ids"`
}
