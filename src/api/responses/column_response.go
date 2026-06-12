package responses

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
)

type ColumnResponse struct {
	ID       uuid.UUID `json:"id"`
	BoardID  uuid.UUID `json:"board_id"`
	Title    string    `json:"title"`
	Position int       `json:"position"`
	Tasks    []TaskResponse `json:"tasks"`
}

func NewColumnResponse(col *models.Column) ColumnResponse {
	columnResponse := ColumnResponse{
		ID:       col.ID,
		Title:    col.Title,
		BoardID:  col.BoardID,
		Position: col.Position,
		Tasks:    []TaskResponse{},
	}

	if len(col.Tasks) > 0 {
		taskResponses := make([]TaskResponse, len(col.Tasks))
		for i, task := range col.Tasks {
			taskResponses[i] = *NewTaskResponse(&task)
		}
		columnResponse.Tasks = taskResponses
	}

	return columnResponse
}
