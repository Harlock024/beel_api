package handlers

import (
	"beel_api/src/dtos"
	"beel_api/src/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshHandler struct {
	service *services.RefreshServices
}

func NewRefreshHandler(service *services.RefreshServices) *RefreshHandler {
	return &RefreshHandler{service: service}
}

func (h *RefreshHandler) RefreshToken(c *gin.Context) {
	refresh_token := dtos.RefreshRequest{}
	if err := c.ShouldBindJSON(&refresh_token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	resposense, err := h.service.Refresh(refresh_token.RefreshToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}
	if resposense == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  resposense.AccessToken,
		"refresh_token": resposense.RefreshToken,
	})
	return
}
