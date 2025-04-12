package handlers

import (
	"beel_api/src/api/utils"
	"beel_api/src/context"
	"beel_api/src/dtos"
	"beel_api/src/models"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
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
				"error":   "user not found",
				"login":   login,
				"user":    userFound,
				"message": "user not found",
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

	token, err := utils.CreateToken(userFound.Username)
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
	newUser.Email = user.Email
	newUser.Username = user.Username
	newUser.Password = passwordHashed

	context.DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"message": "user saved",
	})
}

func GetUser(c *gin.Context) {
	var users []models.User
	context.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	context.DB.Where("id=?", c.Param("id")).Find(&user)

	if user.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
