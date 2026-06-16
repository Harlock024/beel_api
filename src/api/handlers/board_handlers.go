package handlers

import (
	"beel_api/src/dtos"
	"beel_api/src/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BoardHandler struct {
	service *services.BoardService
}

func NewBoardHandler(service *services.BoardService) *BoardHandler {
	return &BoardHandler{service: service}
}

func (h *BoardHandler) GetBoards(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	boards, err := h.service.GetBoardsByUserId(uuid.MustParse(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"boards": boards})
}

func (h *BoardHandler) GetBoard(c *gin.Context) {
	boardID := c.Param("board_id")
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Board ID is required"})
		return
	}

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	board, err := h.service.GetBoardById(uuid.MustParse(boardID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"board": board})
}

func (h *BoardHandler) CreateBoard(c *gin.Context) {
	var dto dtos.BoardDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	board, err := h.service.CreateBoard(uuid.MustParse(userID), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"board": board})
}

func (h *BoardHandler) UpdateBoard(c *gin.Context) {
	boardID := c.Param("board_id")
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Board ID is required"})
		return
	}

	var dto dtos.BoardDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	board, err := h.service.UpdateBoard(uuid.MustParse(boardID), uuid.MustParse(userID), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"board": board})
}

func (h *BoardHandler) DeleteBoard(c *gin.Context) {
	boardID := c.Param("board_id")
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Board ID is required"})
		return
	}

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	if err := h.service.DeleteBoard(uuid.MustParse(boardID), uuid.MustParse(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
