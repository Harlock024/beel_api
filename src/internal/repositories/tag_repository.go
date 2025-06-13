package repositories

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}
func (r *TagRepository) GetAllTagsByUserId(userId uuid.UUID) ([]*models.Tag, error) {
	var tags []*models.Tag
	if err := r.db.Where("user_id = ?", userId).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagRepository) GetTagById(tagId uuid.UUID) (*models.Tag, error) {
	var tag models.Tag
	if err := r.db.Where("id = ?", tagId).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) CreateTag(tag *models.Tag) (*models.Tag, error) {
	if err := r.db.Create(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *TagRepository) UpdateTag(tag *models.Tag) (*models.Tag, error) {
	if err := r.db.Save(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}
func (r *TagRepository) DeleteTag(tag *models.Tag) error {
	if err := r.db.First(tag).Delete(models.Tag{}).Error; err != nil {
		return err
	}
	return nil
}
