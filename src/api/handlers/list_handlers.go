package handlers

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ListHandler struct {
	service *services.ListService
}

func NewListHandler(service *services.ListService) *ListHandler {
	return &ListHandler{service: service}
}
func (h *ListHandler) FindListsByUser(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	listsRes, err := h.service.GetAllListByUserId(uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(listsRes) == 0 {
		c.JSON(http.StatusOK, gin.H{"lists": []responses.ListResponse{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lists": listsRes})
	return

}

func (h *ListHandler) CreateList(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	var list *dtos.ListDTO
	if err := c.BindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	newList, err := h.service.CreateList(list, uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"list": newList})
	return

}

func (h *ListHandler) UpdateList(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	list_id := c.Param("id")
	if list_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "List ID is required"})
		return
	}
	var list *dtos.ListDTO

	if err := c.BindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedList, err := h.service.UpdateList(uuid.MustParse(list_id), *list)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": updatedList})
	return
}

func (h *ListHandler) DeleteList(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	list_id := c.Param("id")
	if list_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "List ID is required"})
		return
	}

	if err := h.service.DeleteList(uuid.MustParse(list_id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
	return
}
