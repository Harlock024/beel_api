package dtos

type ColumnDTO struct {
	Title    string `json:"title" binding:"required"`
	Position *int   `json:"position"`
}

type UpdateColumnDTO struct {
	Title    string `json:"title"`
	Position *int   `json:"position"`
}
