package main

import (
	"beel_api/src/api"
	"beel_api/src/api/handlers"
	"beel_api/src/context"

	"github.com/gin-gonic/gin"
)

func main() {
	context.InitDB()
	r := gin.Default()

	r.POST("/auth/register", handlers.CreateUser)
	r.POST("/auth/login", handlers.LoginHandler)

	r.Use(api.AuthMiddleware())

	r.GET("/users/me", handlers.GetUser)

	r.POST("/api/tasks", handlers.CreateTask)
	r.GET("/api/tasks", handlers.GetUserTasks)

	r.Run("localhost:8080")
}
