package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey" `
	Username string    `json:"username"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password"`
	Tasks    []Task    `json:"tasks" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
