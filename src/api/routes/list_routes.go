package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func ListRoutes(router *gin.RouterGroup) {
	router.GET("/lists", handlers.FindListsByUser)
	router.POST("/lists", handlers.CreateList)
	router.PUT("/lists/:id", handlers.UpdateList)
	router.DELETE("/lists/:id", handlers.DeleteList)
}
