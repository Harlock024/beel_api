package responses

import "github.com/google/uuid"

type ListResponse struct {
	ID    uuid.UUID      `json:"id"`
	Title string         `json:"title"`
	Color string         `json:"color"`
	Tasks []TaskResponse `json:"tasks"`
}
