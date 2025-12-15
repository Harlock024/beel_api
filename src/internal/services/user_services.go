package services

import (
	"beel_api/src/dtos"
	"beel_api/src/internal/repositories"
)

type UserServices struct {
	repo *repositories.UserRepository;
}


func NewUserServices (userRepository *repositories.UserRepository) *UserServices {
	return  &UserServices{repo:userRepository }
}

func (s *UserServices) UpdateUser(dto dtos.UpdateUser){
}	
