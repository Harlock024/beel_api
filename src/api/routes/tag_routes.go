package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func TagRoutes(router *gin.RouterGroup) {
	router.POST("/tags", handlers.CreateTag)
	router.GET("/tags", handlers.GetTags)
	router.GET("/tags/:id", handlers.GetTag)
	router.PUT("/tags/:id", handlers.UpdateTag)
	router.DELETE("/tags/:id", handlers.DeleteTag)
}
