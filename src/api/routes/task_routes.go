package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.RouterGroup, taskHandler *handlers.TaskHandler) {
	router.POST("/lists/:id/tasks", taskHandler.CreateTask)
	router.GET("/lists/:id/tasks", taskHandler.GetTasks)
	router.GET("/tasks/:id", taskHandler.GetTask)
	router.PATCH("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
