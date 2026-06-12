package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"

	"github.com/google/uuid"
)

type BoardService struct {
	repo *repositories.BoardRepository
}

func NewBoardService(repo *repositories.BoardRepository) *BoardService {
	return &BoardService{repo: repo}
}

func (s *BoardService) GetBoardsByUserId(userId uuid.UUID) ([]responses.BoardResponse, error) {
	boards, err := s.repo.GetBoardsByUserId(userId)
	if err != nil {
		return nil, err
	}
	var boardResponses []responses.BoardResponse
	for _, board := range boards {
		boardResponses = append(boardResponses, responses.NewBoardResponseSummary(&board))
	}
	if boardResponses == nil {
		return []responses.BoardResponse{}, nil
	}
	return boardResponses, nil
}

func (s *BoardService) GetBoardById(boardId uuid.UUID) (*responses.BoardResponse, error) {
	board, err := s.repo.GetBoardById(boardId)
	if err != nil {
		return nil, err
	}
	resp := responses.NewBoardResponse(board)
	return &resp, nil
}

func (s *BoardService) CreateBoard(userId uuid.UUID, dto *dtos.BoardDTO) (*responses.BoardResponse, error) {
	board := &models.Board{
		ID:     uuid.New(),
		UserID: userId,
		Title:  dto.Title,
	}

	created, err := s.repo.CreateBoard(board)
	if err != nil {
		return nil, err
	}

	resp := responses.NewBoardResponseSummary(created)
	return &resp, nil
}

func (s *BoardService) UpdateBoard(boardId uuid.UUID, dto *dtos.BoardDTO) (*responses.BoardResponse, error) {
	existing, err := s.repo.GetBoardById(boardId)
	if err != nil {
		return nil, err
	}

	if dto.Title != "" {
		existing.Title = dto.Title
	}

	updated, err := s.repo.UpdateBoard(existing)
	if err != nil {
		return nil, err
	}

	resp := responses.NewBoardResponseSummary(updated)
	return &resp, nil
}

func (s *BoardService) DeleteBoard(boardId uuid.UUID) error {
	return s.repo.DeleteBoard(boardId)
}
