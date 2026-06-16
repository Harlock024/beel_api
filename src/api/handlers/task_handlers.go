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

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task dtos.NewTaskDTO

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

	var listID *uuid.UUID
	if list_id := c.Param("id"); list_id != "" {
		parsed := uuid.MustParse(list_id)
		listID = &parsed
	}

	taskRes, err := h.service.CreateTask(&task, listID, uuid.MustParse(user_id))
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

	h.service.DeleteTask(uuid.MustParse(taskID), uuid.MustParse(user_id))

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

	taskRes, err := h.service.UpdateTask(uuid.MustParse(taskID), uuid.MustParse(user_id), &task)
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

	tasks, err := h.service.GetTasks(filter, uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"tasks": []*responses.TaskResponse{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	return
}

func (h *TaskHandler) GetSubtasks(c *gin.Context) {
	parentID := c.Param("id")
	if parentID == "" {
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

	subtasks, err := h.service.GetSubtasksByParentId(uuid.MustParse(parentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subtasks": subtasks})
}

func (h *TaskHandler) AddSubtask(c *gin.Context) {
	parentID := c.Param("id")
	if parentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	var dto dtos.NewTaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
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

	subtask, err := h.service.AddSubtask(uuid.MustParse(parentID), uuid.MustParse(user_id), &dto, uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"subtask": subtask})
}

func (h *TaskHandler) RemoveSubtask(c *gin.Context) {
	parentID := c.Param("id")
	subtaskID := c.Param("sub_id")
	if parentID == "" || subtaskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID and Subtask ID are required"})
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

	err := h.service.RemoveSubtask(uuid.MustParse(parentID), uuid.MustParse(subtaskID), uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *TaskHandler) GetTasksByTag(c *gin.Context) {
	tagID := c.Query("tag")
	if tagID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag query parameter is required"})
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

	tasks, err := h.service.GetTasksByTag(uuid.MustParse(tagID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *TaskHandler) AddTagToTask(c *gin.Context) {
	taskID := c.Param("id")
	tagID := c.Param("tag_id")
	if taskID == "" || tagID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID and Tag ID are required"})
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

	task, err := h.service.AddTagToTask(uuid.MustParse(taskID), uuid.MustParse(tagID), uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (h *TaskHandler) RemoveTagFromTask(c *gin.Context) {
	taskID := c.Param("id")
	tagID := c.Param("tag_id")
	if taskID == "" || tagID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID and Tag ID are required"})
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

	task, err := h.service.RemoveTagFromTask(uuid.MustParse(taskID), uuid.MustParse(tagID), uuid.MustParse(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (h *TaskHandler) GetTaskCount(c *gin.Context) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claims := claimsRaw.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in claims"})
		return
	}

	filter := c.Query("filter")

	count, err := h.service.GetTaskCount(uuid.MustParse(userID), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

// func (h *TaskHandler) BatchUpdateTasks(c *gin.Context) {
// 	var raw json.RawMessage
// 	if err := c.ShouldBindJSON(&raw); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var dto dtos.BatchUpdateDTO
// 	if err := json.Unmarshal(raw, &dto); err != nil || len(dto.Tasks) == 0 {
// 		var tasksArr []dtos.BatchUpdateTaskDTO
// 		if err := json.Unmarshal(raw, &tasksArr); err != nil || len(tasksArr) == 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 			return
// 		}
// 		dto.Tasks = tasksArr
// 	}

// 	if err := h.service.BatchUpdateTasks(&dto); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Tasks updated successfully"})
// }
