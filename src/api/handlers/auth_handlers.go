package handlers

import (
	"beel_api/src/api/responses"
	"beel_api/src/db"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/services"

	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var user dtos.RegisterDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Invalido"})
		return
	}
	resposense, err := h.service.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	secure := gin.Mode() == gin.ReleaseMode
	c.SetCookie("access_token", resposense.AccessToken, 3600, "/", "", secure, true)
	c.SetCookie("refresh_token", resposense.RefreshToken, 3600*24*30, "/", "", secure, true)

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
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var login dtos.LoginDTO

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Invalido"})
		return
	}
	if login.Email == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	resposense, err := h.service.Login(login)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		return
	}
	// Set the access token as a cookie
	// localhost is used for testing purposes

	secure := gin.Mode() == gin.ReleaseMode
	c.SetCookie("access_token", resposense.AccessToken, 3600, "/", "", secure, true)
	c.SetCookie("refresh_token", resposense.RefreshToken, 3600*24*30, "/", "", secure, true)

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

func GetMe(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	var user models.User
	if err := db.DB.First(&user, "id = ?", user_id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var UserFoundResponse responses.UserResponse
	UserFoundResponse.ID = user.ID
	UserFoundResponse.Username = user.Username
	UserFoundResponse.Email = user.Email
	UserFoundResponse.AvatarURL = user.AvatarURL

	c.JSON(http.StatusOK, gin.H{
		"user": UserFoundResponse,
	})
}

// func LogOut(c *gin.Context) {
// 	claimsRaw, exists := c.Get("claims")
// 	if !exists {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
// 		return
// 	}
// 	claims := claimsRaw.(jwt.MapClaims)
// 	user_id := claims["user_id"].(string)

// 	var refreshTokens models.RefreshToken
// 	err := db.DB.Delete("user_id=?", user_id).Find(&refreshTokens).Error

// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found",
// 			"logged_out": false})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"logged_out": true,
// 	})
// }
