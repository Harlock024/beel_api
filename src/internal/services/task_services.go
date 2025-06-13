package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"

	"github.com/google/uuid"
)

type TaskService struct {
	repo *repositories.TaskRepository
}

func NewTaskService(taskRepository *repositories.TaskRepository) *TaskService {
	return &TaskService{repo: taskRepository}
}

func (s *TaskService) GetTasksByListId(listId uuid.UUID) ([]*responses.TaskResponse, error) {
	tasks, err := s.repo.GetTasksByListId(listId)
	if err != nil {
		return nil, err
	}

	var taskResponses []*responses.TaskResponse
	for _, task := range tasks {
		taskResponse := responses.NewTaskResponse(task)
		taskResponses = append(taskResponses, taskResponse)
	}
	return taskResponses, nil
}

func (s *TaskService) CreateTask(task *dtos.NewTaskDTO) (*responses.TaskResponse, error) {

	newTask := &models.Task{
		ID:    uuid.New(),
		Title: task.Title,
	}
	err := s.repo.CreateTask(newTask)
	if err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(newTask), nil
}

func (s *TaskService) UpdateTask(task_id uuid.UUID, task *dtos.UpdateTaskDTO) (*responses.TaskResponse, error) {

	existingTask, err := s.repo.GetTaskById(task_id)
	if err != nil {
		return nil, err
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Status = task.Status
	existingTask.DueDate = task.DueDate

	err = s.repo.UpdateTask(existingTask)
	if err != nil {
		return nil, err
	}

	return responses.NewTaskResponse(existingTask), nil
}

func (s *TaskService) DeleteTask(task_id uuid.UUID) error {

	existingTask, err := s.repo.GetTaskById(task_id)
	if err != nil {
		return err
	}
	err = s.repo.DeleteTask(existingTask)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetTaskById(task_id uuid.UUID) (*responses.TaskResponse, error) {
	existingTask, err := s.repo.GetTaskById(task_id)
	if err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(existingTask), nil
}
