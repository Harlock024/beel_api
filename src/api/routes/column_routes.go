package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func ColumnRoutes(router *gin.RouterGroup, columnHandler *handlers.ColumnHandler) {
	router.PATCH("/columns/:column_id", columnHandler.UpdateColumn)
	router.DELETE("/columns/:column_id", columnHandler.DeleteColumn)
}
