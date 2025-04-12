package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func NewUser(username string, email string, password string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
		Password: password,
	}
}
