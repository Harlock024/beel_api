package dtos

type BoardDTO struct {
	Title string `json:"title" binding:"required"`
}
