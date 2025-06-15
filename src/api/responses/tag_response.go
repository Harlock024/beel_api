package responses

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
)

type TagResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Color string    `json:"color"`
}

type TagResponses []TagResponse

func NewTagResponse(tag *models.Tag) TagResponse {
	return TagResponse{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}
