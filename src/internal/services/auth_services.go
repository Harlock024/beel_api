package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"
	"beel_api/src/pkg/utils"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo    *repositories.UserRepository
	refreshRepo *repositories.RefreshRepository
}

func NewAuthService(userRepository *repositories.UserRepository, refresh *repositories.RefreshRepository) *AuthService {
	return &AuthService{userRepo: userRepository, refreshRepo: refresh}
}

func (s *AuthService) Register(dto dtos.RegisterDTO) (*responses.LoginResponse, error) {

	hashedPassword, error := utils.HashPassword(dto.Password)
	if error != nil {
		return nil, error
	}

	user := models.User{
		ID:       uuid.New(),
		Email:    dto.Email,
		Username: dto.Username,
		Password: hashedPassword,
	}
	err := s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.Username, user.ID)

	if err != nil {
		return nil, err
	}
	if err := s.refreshRepo.SaveRefreshToken(&models.RefreshToken{
		ID:          uuid.New(),
		UserID:      user.ID,
		HashedToken: refreshToken,
	}); err != nil {
		return nil, err
	}
	loginResponse := &responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: responses.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}
	return loginResponse, nil
}

func (s *AuthService) Login(dto dtos.LoginDTO) (*responses.LoginResponse, error) {

	user, err := s.userRepo.GetUserByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	compare := utils.ComparePassword(user.Password, dto.Password)

	if !compare {
		return nil, err
	}
	accessToken, refreshToken, err := utils.GenerateTokens(user.Username, user.ID)
	if err != nil {
		return nil, err
	}
	loginResponse := &responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: responses.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	if err := s.refreshRepo.SaveRefreshToken(&models.RefreshToken{
		ID:          uuid.New(),
		UserID:      user.ID,
		HashedToken: refreshToken,
	}); err != nil {
		return nil, err
	}
	return loginResponse, nil
}
