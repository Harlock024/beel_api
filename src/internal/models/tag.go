package models

import "github.com/google/uuid"

type Tag struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name   string    `json:"name" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" gorm:"index;not null"`
	Color  string    `json:"color"`
	Tasks  []Task    `gorm:"many2many:task_tags"`
}
