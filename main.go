package main

import (
	"beel_api/src/api"
	"beel_api/src/api/handlers"
	"beel_api/src/context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	context.InitDB()
	r := gin.Default()

	r.POST("/auth/register", handlers.CreateUser)
	r.POST("/auth/login", handlers.LoginHandler)
	r.GET("/users", handlers.GetUser)

	r.Use(api.AuthMiddleware())

	r.GET("/users/me", func(ctx *gin.Context) {
		log.Printf("user me")

	})

	r.Run("localhost:8080")
}
