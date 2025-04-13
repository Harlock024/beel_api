package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", handlers.LoginHandler)
	router.POST("/register", handlers.RegisterHandler)
	router.GET("/me", handlers.GetMe)

}
