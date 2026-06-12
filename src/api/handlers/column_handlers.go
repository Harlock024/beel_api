package handlers

import (
	"beel_api/src/dtos"
	"beel_api/src/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ColumnHandler struct {
	service *services.ColumnService
}

func NewColumnHandler(service *services.ColumnService) *ColumnHandler {
	return &ColumnHandler{service: service}
}

func (h *ColumnHandler) GetColumns(c *gin.Context) {
	boardID := c.Param("board_id")
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Board ID is required"})
		return
	}

	columns, err := h.service.GetColumnsByBoardId(uuid.MustParse(boardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"columns": columns})
}

func (h *ColumnHandler) CreateColumn(c *gin.Context) {
	boardID := c.Param("board_id")
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Board ID is required"})
		return
	}

	var dto dtos.ColumnDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	column, err := h.service.CreateColumn(uuid.MustParse(boardID), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"column": column})
}

func (h *ColumnHandler) UpdateColumn(c *gin.Context) {
	columnID := c.Param("column_id")
	if columnID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Column ID is required"})
		return
	}

	var dto dtos.UpdateColumnDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	column, err := h.service.UpdateColumn(uuid.MustParse(columnID), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"column": column})
}

func (h *ColumnHandler) DeleteColumn(c *gin.Context) {
	columnID := c.Param("column_id")
	if columnID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Column ID is required"})
		return
	}

	if err := h.service.DeleteColumn(uuid.MustParse(columnID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
