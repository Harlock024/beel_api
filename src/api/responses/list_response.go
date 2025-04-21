package responses

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
)

type ListResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Color string    `json:"color"`
}

func NewListResponse(list models.List) ListResponse {
	listResponse := ListResponse{
		ID:    list.ID,
		Title: list.Title,
		Color: list.Color,
	}
	return listResponse
}
