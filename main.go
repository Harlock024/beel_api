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

	// LIST ENDPOINTS

	// create list
	r.POST("/api/lists", handlers.CreateList)

	// list lists
	r.GET("/api/lists", handlers.FindListByUser)

	// update list
	r.PUT("/api/lists/:id", handlers.UpdateList)

	// delete list
	r.DELETE("/api/lists/:id", handlers.DeleteList)

	// TAGS ENDPOINTS

	// list tags
	r.GET("/api/tags", handlers.FindTagsByUser)

	// get tag
	r.GET("/api/tags/:id", handlers.GetTask)

	// create tag
	r.POST("/api/tags", handlers.CreateTag)

	// update tag
	r.PUT("/api/tags/:id", handlers.UpdateTag)

	// delete tag
	r.DELETE("/api/tags/:id", handlers.DeleteTag)

	r.Run("localhost:8080")
}
