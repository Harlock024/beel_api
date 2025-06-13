package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey" `
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	AvatarURL string    `json:"avatar_url"`
	Lists     []List    `gorm:"foreignKey:UserID"`
	Tasks     []Task    `gorm:"foreignKey:UserID"`
	Tags      []Tag     `gorm:"foreignkey:CreatedBy"`
}
