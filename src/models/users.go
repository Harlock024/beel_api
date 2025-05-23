package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey" `
	Username string    `json:"username"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password"`
	List     []List    `gorm:"foreignKey:UserID"`
	Tasks    []Task    `gorm:"foreignKey:UserID"`
	Tags     []Tag     `gorm:"foreignkey:CreatedBy"`
}
