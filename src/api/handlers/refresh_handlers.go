package handlers

import (
	"beel_api/src/dtos"
	"beel_api/src/internal/services"
	"beel_api/src/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

	refresh, err := jwt.Parse(refresh_token.RefreshToken, func(refresh *jwt.Token) (any, error) {
		if _, ok := refresh.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return utils.GetRefreshSecretKey(), nil
	})

	if err != nil || !refresh.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, expired or invalid refresh token"})
		return
	}

	claims, ok := refresh.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}
	user_id, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
		return
	}
	resposense, err := h.service.Refresh(refresh_token.RefreshToken, uuid.MustParse(user_id))

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
