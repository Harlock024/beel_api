package responses

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
)

type TaskResponse struct {
	ID          uuid.UUID       `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	UserID      uuid.UUID       `json:"user_id"`
	Status      string          `json:"status"`
	ListID      uuid.UUID       `json:"list_id"`
	IsCompleted bool            `json:"is_completed"`
	Tags        []TagResponse   `json:"tags"`
	DueDate     string          `json:"due_date"`
	ParentID    *uuid.UUID      `json:"parent_id"`
	Subtasks    []*TaskResponse `json:"subtasks"`
}
type TaskResponses []TaskResponse

func NewTaskResponse(task *models.Task) *TaskResponse {

	var taskResponse TaskResponse
	taskResponse.ID = task.ID
	taskResponse.Title = task.Title
	taskResponse.Description = task.Description
	taskResponse.UserID = task.UserID
	taskResponse.Status = task.Status
	taskResponse.DueDate = task.DueDate
	taskResponse.IsCompleted = task.IsCompleted
	taskResponse.ParentID = task.ParentID

	if task.ListID != uuid.Nil {
		taskResponse.ListID = task.ListID
	} else {
		taskResponse.ListID = uuid.Nil
	}

	if len(task.Tags) > 0 {
		tags := make([]TagResponse, len(task.Tags))
		for i := range task.Tags {
			tags[i] = NewTagResponse(&task.Tags[i])
		}
		taskResponse.Tags = tags
	}

	if len(task.Subtasks) > 0 {
		subtasks := make([]*TaskResponse, len(task.Subtasks))
		for i := range task.Subtasks {
			subtasks[i] = NewTaskResponse(&task.Subtasks[i])
		}
		taskResponse.Subtasks = subtasks
	}

	return &taskResponse
}
