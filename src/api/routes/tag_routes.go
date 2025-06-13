package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func TagRoutes(router *gin.RouterGroup, tagHandler *handlers.TagHandler) {
	router.POST("/tags", tagHandler.CreateTag)
	router.GET("/tags", tagHandler.GetTags)
	// router.GET("/tags/:id", tagHandler.GetTag)
	router.PUT("/tags/:id", tagHandler.UpdateTag)
	router.DELETE("/tags/:id", tagHandler.DeleteTag)
}
