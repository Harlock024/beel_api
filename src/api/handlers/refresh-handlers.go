package handlers

import (
	"beel_api/src/api/responses"
	"beel_api/src/db"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RefreshTokenHandler(c *gin.Context) {
	var refresh_token dtos.RefreshRequest

	if err := c.ShouldBindJSON(&refresh_token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Invalid refresh token payload"})
	}
	hashedToken := utils.HashToken(refresh_token.RefreshToken)

	var savedToken models.RefreshToken

	err := db.DB.Where("hashed_token = ?", hashedToken).First(&savedToken).Error
	if err != nil || savedToken.IsRevoked || time.Now().After(savedToken.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := db.DB.First(&user, savedToken.UserID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	db.DB.Delete(&savedToken)

	accessToken, err := utils.CreateAccessToken(user.Username, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}
	newRefreshToken, err := utils.CreateRefreshToken(user.Username, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create refresh token"})
		return
	}

	newToken := models.RefreshToken{
		HashedToken: utils.HashToken(newRefreshToken),
		UserID:      user.ID,
		ExpiresAt:   time.Now().Add(7 * 24 * 30),
		IsRevoked:   false,
	}
	if err := db.DB.Create(&newToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store refresh token"})
		return
	}

	var UserFoundResponse responses.UserResponse
	UserFoundResponse.ID = user.ID
	UserFoundResponse.Username = user.Username
	UserFoundResponse.Email = user.Email
	UserFoundResponse.AvatarURL = user.AvatarURL

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": newRefreshToken,
		"user":         UserFoundResponse,
	})
}
