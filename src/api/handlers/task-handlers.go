package handlers

import (
	"beel_api/src/context"
	"beel_api/src/dtos"
	"beel_api/src/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateTask(c *gin.Context) {
	var task dtos.NewTaskDTO

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newTask models.Task
	newTask.ID = uuid.New()
	newTask.Title = task.Title
	newTask.UserID = uuid.MustParse(user_id)
	newTask.Status = false

	err := context.DB.Create(&newTask).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"task": newTask})
	return
}

func GetUserTasks(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	var tasks []models.Task
	err := context.DB.Where("user_id = ?", user_id).Find(&tasks).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	return
}
