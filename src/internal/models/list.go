package models

import "github.com/google/uuid"

type List struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Title  string    `json:"title" gorm:"not null"`
	Color  string    `json:"color" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" gorm:"index;foreignKey:UserID;constraint:OnDelete:CASCADE"`
	User   User      `gorm:"foreignKey:UserID"`
	Tasks  []Task    `gorm:"constraint:OnDelete:CASCADE;"`
}
