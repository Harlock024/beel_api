package handlers

import (
	"beel_api/src/dtos"
	"beel_api/src/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *services.UserServices
}

func NewUserHandler(service *services.UserServices) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	var dto dtos.UpdateUser
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UpdateUser(uuid.MustParse(userID), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
