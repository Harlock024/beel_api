package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	router.PATCH("/profile", userHandler.UpdateProfile)
}
