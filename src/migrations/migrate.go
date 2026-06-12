package migrations

import (
	"beel_api/src/db"
	"beel_api/src/internal/models"

	"log"
)

func Run() {
	err := db.DB.AutoMigrate(&models.User{}, &models.Board{}, &models.Column{}, &models.Task{}, &models.List{}, &models.Tag{}, &models.RefreshToken{}) 

	if err != nil {
		log.Fatal("Error AutoMigrate",err)
	}
}
