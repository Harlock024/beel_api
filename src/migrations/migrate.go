package migrations

import (
	"beel_api/src/db"
	"beel_api/src/internal/models"
)

func Run() {
	err := db.DB.AutoMigrate(&models.User{}, &models.Task{}, &models.List{}, &models.Tag{}, &models.RefreshToken{})

	if err != nil {
		panic("Error during migration: " + err.Error())
	}
}
