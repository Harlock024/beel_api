package dtos

import "github.com/google/uuid"

type BatchUpdateDTO struct {
	Tasks []BatchUpdateTaskDTO `json:"tasks" binding:"required"`
}

type BatchUpdateTaskDTO struct {
	ID       uuid.UUID  `json:"id" binding:"required"`
	ColumnID *uuid.UUID `json:"column_id"`
	Position *int       `json:"position"`
}
