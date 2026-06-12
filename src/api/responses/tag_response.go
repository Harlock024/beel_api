package responses

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
)

type TagResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	TaskCount int       `json:"task_count"`
}

type TagResponses []TagResponse

func NewTagResponse(tag *models.Tag) TagResponse {
	return TagResponse{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}

func NewTagResponseWithCount(tag *models.Tag, count int) TagResponse {
	return TagResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		Color:     tag.Color,
		TaskCount: count,
	}
}
