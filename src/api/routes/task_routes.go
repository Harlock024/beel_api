package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.RouterGroup) {
	router.POST("/lists/:id/tasks", handlers.CreateTask)
	router.GET("/lists/:id/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
}
