package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"

	"github.com/google/uuid"
)

type ColumnService struct {
	repo *repositories.ColumnRepository
}

func NewColumnService(repo *repositories.ColumnRepository) *ColumnService {
	return &ColumnService{repo: repo}
}

func (s *ColumnService) GetColumnsByBoardId(boardId uuid.UUID) ([]responses.ColumnResponse, error) {
	columns, err := s.repo.GetColumnsByBoardId(boardId)
	if err != nil {
		return nil, err
	}
	var columnResponses []responses.ColumnResponse
	for _, col := range columns {
		columnResponses = append(columnResponses, responses.NewColumnResponse(&col))
	}
	if columnResponses == nil {
		return []responses.ColumnResponse{}, nil
	}
	return columnResponses, nil
}

func (s *ColumnService) CreateColumn(boardId uuid.UUID, dto *dtos.ColumnDTO) (*responses.ColumnResponse, error) {
	maxPos, err := s.repo.GetMaxPosition(boardId)
	if err != nil {
		return nil, err
	}

	column := &models.Column{
		ID:       uuid.New(),
		BoardID:  boardId,
		Title:    dto.Title,
		Position: maxPos + 1,
	}

	created, err := s.repo.CreateColumn(column)
	if err != nil {
		return nil, err
	}

	resp := responses.NewColumnResponse(created)
	return &resp, nil
}

func (s *ColumnService) UpdateColumn(columnId uuid.UUID, dto *dtos.UpdateColumnDTO) (*responses.ColumnResponse, error) {
	existing, err := s.repo.GetColumnById(columnId)
	if err != nil {
		return nil, err
	}

	if dto.Title != "" {
		existing.Title = dto.Title
	}
	if dto.Position != nil {
		existing.Position = *dto.Position
	}

	updated, err := s.repo.UpdateColumn(existing)
	if err != nil {
		return nil, err
	}

	resp := responses.NewColumnResponse(updated)
	return &resp, nil
}

func (s *ColumnService) DeleteColumn(columnId uuid.UUID) error {
	return s.repo.DeleteColumn(columnId)
}
