package responses

import "beel_api/src/internal/models"

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Status      bool   `json:"status"`
	ListID      string `json:"list_id"`
}

type TaskResponses []TaskResponse

func TaskRes(task models.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		UserID:      task.UserID.String(),
		Status:      task.Status,
		ListID:      task.ListID.String(),
	}

}
