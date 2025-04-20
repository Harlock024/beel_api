package responses

import (
	"beel_api/src/internal/models"
)

type TaskResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	UserID      string   `json:"user_id"`
	Status      bool     `json:"status"`
	ListID      string   `json:"list_id"`
	Tags        []string `json:"tags"`
}

type TaskResponses []TaskResponse

func TaskRes(task models.Task) TaskResponse {

	var taskResponse TaskResponse
	taskResponse.ID = task.ID.String()
	taskResponse.Title = task.Title
	taskResponse.Description = task.Description
	taskResponse.UserID = task.UserID.String()
	taskResponse.Status = task.Status

	if task.ListID != nil {
		taskResponse.ListID = task.ListID.String()

	} else {
		taskResponse.ListID = ""
	}
	return taskResponse
}
