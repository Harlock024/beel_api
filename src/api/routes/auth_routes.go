package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	router.POST("/login", authHandler.LoginHandler)
	router.POST("/register", authHandler.RegisterHandler)
	// router.GET("/me", handlers.GetMe)

}
