package handlers

import (
	"beel_api/src/api/responses"
	"beel_api/src/db"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"time"

	"beel_api/src/pkg/utils"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func LoginHandler(c *gin.Context) {
	var login dtos.LoginDTO

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	err := db.DB.Where("email=?", login.Email).First(&userFound).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var compare bool
	if compare = utils.ComparePassword(userFound.Password, login.Password); !compare {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password invalid"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(userFound.Username, userFound.ID)

	db.DB.Save(models.RefreshToken{
		ID:          uuid.New(),
		UserID:      userFound.ID,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 30),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		HashedToken: utils.HashToken(refreshToken),
		IsRevoked:   false,
	})

	var UserFoundResponse responses.UserResponse
	UserFoundResponse.ID = userFound.ID
	UserFoundResponse.Username = userFound.Username
	UserFoundResponse.Email = userFound.Email
	UserFoundResponse.AvatarURL = userFound.AvatarURL

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          UserFoundResponse,
	})
}

func RegisterHandler(c *gin.Context) {
	var user dtos.RegisterDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Invalido"})
		return
	}

	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var newUser models.User
	newUser.ID = uuid.New()
	newUser.Email = user.Email
	newUser.Username = user.Username
	newUser.Password = passwordHashed

	err = db.DB.Create(&newUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(newUser.Username, newUser.ID)

	db.DB.Save(models.RefreshToken{
		ID:          uuid.New(),
		UserID:      newUser.ID,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 30),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		HashedToken: utils.HashToken(refreshToken),
		IsRevoked:   false,
	})

	var UserFoundResponse responses.UserResponse
	UserFoundResponse.ID = newUser.ID
	UserFoundResponse.Username = newUser.Username
	UserFoundResponse.Email = newUser.Email
	UserFoundResponse.AvatarURL = newUser.AvatarURL

	c.JSON(http.StatusOK, gin.H{

		"token":         accessToken,
		"refresh_token": refreshToken,
		"user":          UserFoundResponse,
	})
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

func LogOut(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	var refreshTokens models.RefreshToken
	err := db.DB.Delete("user_id=?", user_id).Find(&refreshTokens).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found",
			"logged_out": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"logged_out": true,
	})
}
