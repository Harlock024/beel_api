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
	Status      bool        `json:"status"`
	DueDate     string      `json:"due_date"`
	ListID      *uuid.UUID  `json:"list_id"`
	TagIDs      []uuid.UUID `json:"tag_ids"`
}
