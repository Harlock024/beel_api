package repositories

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *BoardRepository {
	return &BoardRepository{db: db}
}

func (r *BoardRepository) GetBoardsByUserId(userId uuid.UUID) ([]models.Board, error) {
	var boards []models.Board
	if err := r.db.Where("user_id = ?", userId).Order("created_at DESC").Find(&boards).Error; err != nil {
		return nil, err
	}
	return boards, nil
}

func (r *BoardRepository) GetBoardById(id uuid.UUID) (*models.Board, error) {
	var board models.Board
	if err := r.db.Where("id = ?", id).First(&board).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

func (r *BoardRepository) CreateBoard(board *models.Board) (*models.Board, error) {
	if err := r.db.Create(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (r *BoardRepository) UpdateBoard(board *models.Board) (*models.Board, error) {
	if err := r.db.Save(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (r *BoardRepository) DeleteBoard(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Board{}).Error; err != nil {
		return err
	}
	return nil
}
