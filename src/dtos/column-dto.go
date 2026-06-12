package dtos

type ColumnDTO struct {
	Title    string `json:"title" binding:"required"`
	Position *int   `json:"position"`
}
