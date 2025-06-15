package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"

	"github.com/google/uuid"
)

type TagService struct {
	repo *repositories.TagRepository
}

func NewTagService(repo *repositories.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) GetAllTagsByUserId(userId uuid.UUID) ([]*responses.TagResponse, error) {
	tags, err := s.repo.GetAllTagsByUserId(userId)
	if err != nil {
		return nil, err
	}
	var tagResponses []*responses.TagResponse
	for _, tag := range tags {
		tagResponse := responses.NewTagResponse(tag)
		tagResponses = append(tagResponses, &tagResponse)
	}
	return tagResponses, nil
}
func (s *TagService) GetTagById(tagId uuid.UUID) (*responses.TagResponse, error) {
	tag, err := s.repo.GetTagById(tagId)
	if err != nil {
		return nil, err
	}
	tagResponse := responses.NewTagResponse(tag)
	return &tagResponse, nil
}
func (s *TagService) CreateTag(tag *dtos.TagDTO, userId uuid.UUID) (*responses.TagResponse, error) {
	newTag, err := s.repo.CreateTag(&models.Tag{
		ID:     uuid.New(),
		Name:   tag.Name,
		Color:  tag.Color,
		UserID: userId,
	})
	if err != nil {
		return nil, err
	}
	tagResponse := responses.NewTagResponse(newTag)
	return &tagResponse, nil
}
func (s *TagService) UpdateTag(tagId uuid.UUID, tag dtos.TagDTO) (*responses.TagResponse, error) {
	existingTag, err := s.repo.GetTagById(tagId)
	if err != nil {
		return nil, err
	}
	existingTag.Name = tag.Name
	existingTag.Color = tag.Color

	updatedTag, err := s.repo.UpdateTag(existingTag)
	if err != nil {
		return nil, err
	}
	tagResponse := responses.NewTagResponse(updatedTag)
	return &tagResponse, nil
}

func (s *TagService) DeleteTag(tagId uuid.UUID) error {
	tag, err := s.repo.GetTagById(tagId)
	if err != nil {
		return err
	}
	if err := s.repo.DeleteTag(tag); err != nil {
		return err
	}
	return nil
}
