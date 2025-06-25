package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"
	"beel_api/src/pkg/utils"

	"github.com/google/uuid"
)

type RefreshServices struct {
	refreshRepo *repositories.RefreshRepository
	userRepo    *repositories.UserRepository
}

func NewRefreshServices(repo *repositories.RefreshRepository, userRepo *repositories.UserRepository) *RefreshServices {
	return &RefreshServices{
		refreshRepo: repo,
		userRepo:    userRepo,
	}
}

func (s *RefreshServices) Refresh(refresh_token string, user_id uuid.UUID) (*responses.LoginResponse, error) {
	refreshToken, err := s.refreshRepo.FindByRefreshToken(utils.HashToken(refresh_token))
	if err != nil {
		return nil, err
	}

	// Delete the old refresh token
	if err := s.refreshRepo.DeleteRefreshToken(refreshToken); err != nil {
		return nil, err
	}
	// Find the user by ID
	user, err := s.userRepo.GetUserByID(refreshToken.UserID)
	if err != nil {
		return nil, err
	}

	// Generate new tokens (access and refresh)
	newAccess, newRefresh, err := utils.GenerateTokens(user.Username, user.ID)

	if err != nil {
		return nil, err
	}
	// Save the new refresh token hashed
	if err := s.refreshRepo.SaveRefreshToken(&models.RefreshToken{
		ID:          uuid.New(),
		HashedToken: utils.HashToken(newRefresh),
		UserID:      user_id,
		IsRevoked:   false,
	}); err != nil {
		return nil, err
	}

	// Return the new tokens and user information
	return &responses.LoginResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		User: responses.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			AvatarURL: user.AvatarURL,
		},
	}, nil
}
