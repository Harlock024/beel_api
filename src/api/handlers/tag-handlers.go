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

func FindTagsByUser(c *gin.Context) {

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	tags := []models.Tag{}
	err := context.DB.Where("created_by = ?", userID).Find(&tags).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func getTag(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	tag := models.Tag{}
	err := context.DB.First(&tag, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

func CreateTag(c *gin.Context) {
	var tag dtos.TagDTO

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newTag models.Tag
	newTag.ID = uuid.New()
	newTag.Name = tag.Name
	newTag.Color = tag.Color
	newTag.CreatedBy = uuid.MustParse(userID)
	err := context.DB.Create(&newTag).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tag)
}

func UpdateTag(c *gin.Context) {
	var updateTag dtos.TagDTO

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := c.ShouldBindJSON(&updateTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tag models.Tag
	err := context.DB.Where("id = ? AND created_by = ?", taskID, userID).First(&tag).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag not found"})
		return
	}

	tag.Name = updateTag.Name
	tag.Color = updateTag.Color

	err = context.DB.Save(&tag).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

func DeleteTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := context.DB.Delete(&tag).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
