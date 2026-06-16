package repositories

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ColumnRepository struct {
	db *gorm.DB
}

func NewColumnRepository(db *gorm.DB) *ColumnRepository {
	return &ColumnRepository{db: db}
}

func (r *ColumnRepository) GetColumnsByBoardId(boardId uuid.UUID) ([]models.Column, error) {
	var columns []models.Column
	if err := r.db.Preload("Tasks").Where("board_id = ?", boardId).Order("position ASC").Find(&columns).Error; err != nil {
		return nil, err
	}
	return columns, nil
}

func (r *ColumnRepository) GetColumnById(id uuid.UUID) (*models.Column, error) {
	var column models.Column
	if err := r.db.Preload("Tasks").Where("id = ?", id).First(&column).Error; err != nil {
		return nil, err
	}
	return &column, nil
}

func (r *ColumnRepository) GetColumnByIdWithBoard(id uuid.UUID) (*models.Column, error) {
	var column models.Column
	if err := r.db.Preload("Board").Where("id = ?", id).First(&column).Error; err != nil {
		return nil, err
	}
	return &column, nil
}

func (r *ColumnRepository) CreateColumn(column *models.Column) (*models.Column, error) {
	if err := r.db.Create(column).Error; err != nil {
		return nil, err
	}
	return column, nil
}

func (r *ColumnRepository) UpdateColumn(column *models.Column) (*models.Column, error) {
	if err := r.db.Save(column).Error; err != nil {
		return nil, err
	}
	return column, nil
}

func (r *ColumnRepository) DeleteColumn(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Column{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *ColumnRepository) GetMaxPosition(boardId uuid.UUID) (int, error) {
	var result struct {
		MaxPos int
	}
	if err := r.db.Model(&models.Column{}).Select("COALESCE(MAX(position), -1) as max_pos").Where("board_id = ?", boardId).Scan(&result).Error; err != nil {
		return -1, err
	}
	return result.MaxPos, nil
}

func (r *ColumnRepository) ReorderColumns(columns []models.Column) error {
	return r.db.Save(&columns).Error
}
