package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"

	"github.com/google/uuid"
)

type ListService struct {
	repo repositories.ListRepository
}

func NewListService(repo repositories.ListRepository) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) GetAllListByUserId(user_id uuid.UUID) ([]*responses.ListResponse, error) {
	lists, err := s.repo.GetAllListByUserId(user_id)
	if err != nil {
		return nil, err
	}
	var listResponses []*responses.ListResponse
	for _, list := range lists {
		listResponse := responses.NewListResponse(list)
		listResponses = append(listResponses, &listResponse)
	}

	if len(listResponses) == 0 {
		return listResponses, nil
	}
	return listResponses, nil
}

func (s *ListService) GetListById(list_id uuid.UUID) (*responses.ListResponse, error) {
	list, err := s.repo.GetListById(list_id)
	if err != nil {
		return nil, err
	}
	listResponse := responses.NewListResponse(*list)
	return &listResponse, nil
}

func (s *ListService) CreateList(list *dtos.ListDTO, user_id uuid.UUID) (*responses.ListResponse, error) {
	newList, err := s.repo.CreateList(&models.List{
		ID:     uuid.New(),
		Title:  list.Title,
		Color:  list.Color,
		UserID: user_id,
	})
	if err != nil {
		return nil, err
	}
	listResponse := responses.NewListResponse(*newList)
	return &listResponse, nil
}

func (s *ListService) UpdateList(list_id uuid.UUID, list dtos.ListDTO) (*responses.ListResponse, error) {
	existingList, err := s.repo.GetListById(list_id)
	if err != nil {
		return nil, err
	}
	existingList.Title = list.Title
	existingList.Color = list.Color
	updatedList, err := s.repo.UpdateList(existingList)
	if err != nil {
		return nil, err
	}
	listResponse := responses.NewListResponse(*updatedList)
	return &listResponse, nil
}

func (s *ListService) DeleteList(list_id uuid.UUID) error {
	if err := s.repo.DeleteList(list_id); err != nil {
		return err
	}
	return nil
}
