package responses

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
)

type TaskResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id"`
	Status      string    `json:"status"`
	ListID      uuid.UUID `json:"list_id"`
	Tags        []string  `json:"tags"`
	DueDate     string    `json:"due_date"`
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

	if task.ListID != uuid.Nil {
		taskResponse.ListID = task.ListID

	} else {
		taskResponse.ListID = uuid.Nil
	}
	return &taskResponse
}
