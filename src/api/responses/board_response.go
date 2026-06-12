package responses

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
)

type BoardResponse struct {
	ID        uuid.UUID        `json:"id"`
	UserID    uuid.UUID        `json:"user_id"`
	Title     string           `json:"title"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
	Columns   []ColumnResponse `json:"columns"`
}

func NewBoardResponse(board *models.Board) BoardResponse {
	var columns []ColumnResponse
	if len(board.Columns) > 0 {
		columns = make([]ColumnResponse, len(board.Columns))
		for i, col := range board.Columns {
			columns[i] = ColumnResponse{
				ID:       col.ID,
				BoardID:  col.BoardID,
				Title:    col.Title,
				Position: col.Position,
			}
		}
	} else {
		columns = []ColumnResponse{}
	}

	return BoardResponse{
		ID:        board.ID,
		UserID:    board.UserID,
		Title:     board.Title,
		CreatedAt: board.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: board.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		Columns:   columns,
	}
}

func NewBoardResponseSummary(board *models.Board) BoardResponse {
	return BoardResponse{
		ID:        board.ID,
		UserID:    board.UserID,
		Title:     board.Title,
		CreatedAt: board.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: board.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		Columns:   []ColumnResponse{},
	}
}
