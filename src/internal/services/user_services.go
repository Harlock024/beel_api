package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/repositories"

	"github.com/google/uuid"
)

type UserServices struct {
	repo *repositories.UserRepository
}

func NewUserServices(userRepository *repositories.UserRepository) *UserServices {
	return &UserServices{repo: userRepository}
}

func (s *UserServices) UpdateUser(userId uuid.UUID, dto dtos.UpdateUser) (*responses.UserResponse, error) {
	user, err := s.repo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	if dto.Username != "" {
		user.Username = dto.Username
	}
	if dto.AvatarUrl != "" {
		user.AvatarURL = dto.AvatarUrl
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}

	return &responses.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}, nil
}
