package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.RouterGroup, taskHandler *handlers.TaskHandler) {
	router.POST("/lists/:id/tasks", taskHandler.CreateTask)
	router.GET("/lists/:id/tasks", taskHandler.GetTasks)
	router.GET("/tasks", taskHandler.GetTasksByFilter)
	router.GET("/tasks/by-tag", taskHandler.GetTasksByTag)
	router.GET("/tasks/:id", taskHandler.GetTask)
	router.PATCH("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	router.GET("/tasks/:id/subtasks", taskHandler.GetSubtasks)
	router.POST("/tasks/:id/subtasks", taskHandler.AddSubtask)
	router.DELETE("/tasks/:id/subtasks/:sub_id", taskHandler.RemoveSubtask)

	router.POST("/tasks/:id/tags/:tag_id", taskHandler.AddTagToTask)
	router.DELETE("/tasks/:id/tags/:tag_id", taskHandler.RemoveTagFromTask)
}
