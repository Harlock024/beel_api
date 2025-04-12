package handlers

import (
	"beel_api/src/api/utils"
	"beel_api/src/context"
	"beel_api/src/dtos"
	"beel_api/src/models"
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
	err := context.DB.Where("email=?", login.Email).First(&userFound).Error

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

	token, err := utils.CreateToken(userFound.Username, userFound.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func CreateUser(c *gin.Context) {
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

	err = context.DB.Create(&newUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.CreateToken(newUser.Username, newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user saved",
		"token":   token,
	})
}

func GetUser(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	var user models.User
	if err := context.DB.First(&user, "id = ?", user_id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
