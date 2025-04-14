package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func RefreshRoutes(router *gin.RouterGroup) {
	router.GET("/refresh", handlers.RefreshTokenHandler)
}
