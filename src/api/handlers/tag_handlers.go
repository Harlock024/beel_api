package handlers

import (
	"beel_api/src/dtos"
	"beel_api/src/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TagHandler struct {
	service *services.TagService
}

func NewTagHandler(service *services.TagService) *TagHandler {
	return &TagHandler{service: service}
}

func (h *TagHandler) GetTags(c *gin.Context) {

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}
	tagsRes, err := h.service.GetAllTagsByUserId(uuid.MustParse(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tags": tagsRes})

}
func (h *TagHandler) CreateTag(c *gin.Context) {
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

	newTagRes, err := h.service.CreateTag(&tag, uuid.MustParse(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"tag": newTagRes})
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	var tag dtos.TagDTO

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tagId := c.Param("id")
	if tagId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag ID is required"})
		return
	}
	updatedTagRes, err := h.service.UpdateTag(uuid.MustParse(tagId), tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tag": updatedTagRes})

}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}
	tagId := c.Param("id")
	if tagId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag ID is required"})
		return
	}

	err := h.service.DeleteTag(uuid.MustParse(tagId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
	return
}
