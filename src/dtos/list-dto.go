package dtos

type ListDTO struct {
	Title string `json:"title" binding:"required"`
	Color string `json:"color" binding:"required"`
}
