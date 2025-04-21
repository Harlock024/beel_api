package handlers

import (
	"beel_api/src/api/responses"
	"beel_api/src/db"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func FindListsByUser(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	var list []models.List

	if err := db.DB.Where("user_id = ?", user_id).Find(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ListResponses := make([]responses.ListResponse, len(list))
	for i, l := range list {
		ListResponses[i] = responses.NewListResponse(l)
	}
	c.JSON(http.StatusOK, gin.H{"lists": ListResponses})
	return
}

func CreateList(c *gin.Context) {
	var list dtos.ListDTO
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if err := c.BindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newList models.List
	newList.ID = uuid.New()
	newList.UserID = user_id
	newList.Title = list.Name
	newList.Color = list.Color

	if err := db.DB.Create(&newList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"list": responses.NewListResponse(newList)})
	return
}

func DeleteList(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	var list models.List

	if err := db.DB.Where("user_id = ? AND id = ?", user_id, c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Delete(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
	return
}

func UpdateList(c *gin.Context) {
	var list dtos.ListDTO
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if err := db.DB.Where("user_id = ? AND id = ?", user_id, c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedList models.List
	updatedList.Title = list.Name
	updatedList.Color = list.Color

	if err := db.DB.Save(&updatedList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": responses.NewListResponse(updatedList)})
	return
}
