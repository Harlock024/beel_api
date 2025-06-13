package repositories

import (
	"beel_api/src/internal/models"

	"gorm.io/gorm"
)

type RefreshRepository struct {
	db *gorm.DB
}

func NewRefreshRepository(db *gorm.DB) *RefreshRepository {
	return &RefreshRepository{db: db}
}

func (r *RefreshRepository) SaveRefreshToken(refresh *models.RefreshToken) error {
	if err := r.db.Create(refresh).Error; err != nil {
		return err
	}
	return nil
}
func (r *RefreshRepository) FindByRefreshToken(refreshTokenHashed string) (*models.RefreshToken, error) {
	refreshToken := &models.RefreshToken{}
	if err := r.db.Where("HashedToken = ?", refreshTokenHashed).First(refreshToken).Error; err != nil {
		return nil, err
	}
	return refreshToken, nil
}
func (r *RefreshRepository) DeleteRefreshToken(refreshToken *models.RefreshToken) error {
	if err := r.db.Delete(refreshToken).Error; err != nil {
		return err
	}
	return nil
}

func (r *RefreshRepository) DeleteAllRefreshTokensByUserId(userId string) error {
	if err := r.db.Where("user_id = ?", userId).Delete(&models.RefreshToken{}).Error; err != nil {
		return err
	}
	return nil
}
