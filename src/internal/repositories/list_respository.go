package repositories

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) GetAllListByUserId(user_id uuid.UUID) ([]models.List, error) {
	var lists []models.List
	if err := r.db.Where("user_id = ? ", user_id).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *ListRepository) GetListById(list_id uuid.UUID) (*models.List, error) {
	var list models.List
	if err := r.db.Where("id = ?", list_id).Take(&list).Error; err != nil {
		return nil, err
	}
	return &list, nil
}
func (r *ListRepository) CreateList(list *models.List) (*models.List, error) {
	if err := r.db.Create(list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ListRepository) UpdateList(list *models.List) (*models.List, error) {
	if err := r.db.Save(list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ListRepository) DeleteList(list_id uuid.UUID) error {
	if err := r.db.Where("id = ?", list_id).Delete(&models.List{}).Error; err != nil {
		return err
	}
	return nil
}
