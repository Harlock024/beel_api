package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func ListRoutes(router *gin.RouterGroup, listHandler *handlers.ListHandler) {
	router.GET("/lists", listHandler.FindListsByUser)
	router.POST("/lists", listHandler.CreateList)
	router.PUT("/lists/:id", listHandler.UpdateList)
	router.DELETE("/lists/:id", listHandler.DeleteList)
}
