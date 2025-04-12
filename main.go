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

	// AUTH ENDPOINTS

	// register user
	r.POST("/auth/register", handlers.CreateUser)

	// login user
	r.POST("/auth/login", handlers.LoginHandler)

	// logout user
	// todo

	r.Use(api.AuthMiddleware())

	r.GET("/users/me", handlers.GetUser)

	// TASK ENDPOINTS

	// create task
	r.POST("/api/tasks", handlers.CreateTask)

	// list tasks
	r.GET("/api/tasks", handlers.GetUserTasks)

	// update task
	r.PUT("/api/tasks/:id", handlers.UpdateTask)

	// delete task
	r.DELETE("/api/tasks/:id", handlers.DeleteTask)

	r.Run("localhost:8080")
}
