package routes

import (
	"beel_api/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func BoardRoutes(router *gin.RouterGroup, boardHandler *handlers.BoardHandler, columnHandler *handlers.ColumnHandler) {
	boards := router.Group("/boards")
	{
		boards.GET("", boardHandler.GetBoards)
		boards.POST("", boardHandler.CreateBoard)
		boards.GET("/:board_id", boardHandler.GetBoard)
		boards.PATCH("/:board_id", boardHandler.UpdateBoard)
		boards.DELETE("/:board_id", boardHandler.DeleteBoard)

		boards.GET("/:board_id/columns", columnHandler.GetColumns)
		boards.POST("/:board_id/columns", columnHandler.CreateColumn)
	}
}
