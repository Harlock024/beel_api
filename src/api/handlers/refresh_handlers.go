package handlers

import (
	"beel_api/src/api/responses"
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
	var refresh_token dtos.RefreshRequest

	if err := c.ShouldBindJSON(&refresh_token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Invalid refresh token payload"})
	}

	resposense, err := h.service.Refresh(refresh_token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}

	secure := gin.Mode() == gin.ReleaseMode
	c.SetCookie("access_token", resposense.AccessToken, 3600, "/", "localhost", secure, true)

	c.JSON(http.StatusOK, gin.H{
		"refresh_token": resposense.RefreshToken,
		"user": responses.UserResponse{
			ID:        resposense.User.ID,
			Username:  resposense.User.Username,
			Email:     resposense.User.Email,
			AvatarURL: resposense.User.AvatarURL,
		},
	})
	return
}
