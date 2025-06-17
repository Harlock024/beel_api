package handlers

import (
	"beel_api/src/api/responses"
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
	refresh_token, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	resposense, err := h.service.Refresh(refresh_token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}
	if resposense == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	c.SetCookie("access_token", resposense.AccessToken, 3600, "/", "localhost", false, true)
	c.SetCookie("refresh_token", resposense.RefreshToken, 3600*24*7, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user": responses.UserResponse{
			ID:        resposense.User.ID,
			Username:  resposense.User.Username,
			Email:     resposense.User.Email,
			AvatarURL: resposense.User.AvatarURL,
		},
	})
	return
}
