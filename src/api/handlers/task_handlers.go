package handlers

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/services"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task dtos.NewTaskDTO

	var list_id = c.Param("id")

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claimsRaw, exists := c.Get("claims")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	taskRes, err := h.service.CreateTask(&task, uuid.MustParse(list_id), uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"task": taskRes})
	return
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	listID := c.Param("id")
	if listID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "List ID is required"})
		return
	}

	claimsRaw, exists := c.Get("claims")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	tasks, err := h.service.GetTasksByListId(uuid.MustParse(listID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	return
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	task, err := h.service.GetTaskById(uuid.MustParse(taskID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
	return
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	h.service.DeleteTask(uuid.MustParse(taskID))

	c.JSON(http.StatusNoContent, gin.H{"message": "Task deleted successfully"})
	return
}
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	var task dtos.UpdateTaskDTO
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	taskRes, err := h.service.UpdateTask(uuid.MustParse(taskID), &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": taskRes})
	return
}

func (h *TaskHandler) GetTasksByFilter(c *gin.Context) {
	filter := c.Query("filter")
	if filter == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filter is required"})
		return
	}
	fmt.Println("Filter:", filter)

	tasks, err := h.service.GetTasksByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"tasks": []responses.TaskResponses{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	return
}
