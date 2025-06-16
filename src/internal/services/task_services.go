package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"
	"time"

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
	if len(taskResponses) == 0 {
		return []*responses.TaskResponse{}, nil
	}
	return taskResponses, nil
}

func (s *TaskService) CreateTask(task *dtos.NewTaskDTO, list_id uuid.UUID, user_id uuid.UUID) (*responses.TaskResponse, error) {

	newTask := &models.Task{
		ID:     uuid.New(),
		Title:  task.Title,
		ListID: list_id,
		UserID: user_id,
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
	if task.Title != "" {
		existingTask.Title = task.Title
	}
	if task.Description != "" {
		existingTask.Description = task.Description
	}
	if task.Status != "" {
		existingTask.Status = task.Status
	}
	if task.DueDate != "" {
		existingTask.DueDate = task.DueDate
	}
	if task.ListID != nil {
		existingTask.ListID = *task.ListID
	}
	if task.IsCompleted {
		existingTask.IsCompleted = task.IsCompleted
	}
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

func (s TaskService) GetTasksByFilter(filter string) ([]*responses.TaskResponse, error) {
	if filter == "" {
		return nil, nil
	}

	loc, _ := time.LoadLocation("America/Mexico_City")

	if filter == "today" {
		now := time.Now().In(loc)
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).UTC()
		end := start.Add(24 * time.Hour).UTC()
		tasks, err := s.repo.GetTasksByFilter(start.String(), end.String())
		if err != nil {
			return nil, err
		}

		if len(tasks) == 0 {
			return nil, nil
		}
		var taskResponses []*responses.TaskResponse
		for _, task := range tasks {
			taskResponse := responses.NewTaskResponse(task)
			taskResponses = append(taskResponses, taskResponse)
		}

		return taskResponses, nil
	} else if filter == "upcoming" {
		start := time.Now().In(loc).Truncate(24 * time.Hour).Add(24 * time.Hour).UTC()
		end := start.Add(7 * 24 * time.Hour).UTC()
		tasks, err := s.repo.GetTasksByFilter(start.String(), end.String())
		if err != nil {
			return nil, err
		}

		if len(tasks) == 0 {
			return nil, nil
		}
		var taskResponses []*responses.TaskResponse
		for _, task := range tasks {
			taskResponse := responses.NewTaskResponse(task)
			taskResponses = append(taskResponses, taskResponse)
		}
		return taskResponses, nil
	}
	// } else if filter == "overdue" {
	// 	tasks, err := s.repo.GetTasksByFilter("overdue")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if len(tasks) == 0 {
	// 		return nil, nil
	// 	}
	// 	var taskResponses []*responses.TaskResponse
	// 	for _, task := range tasks {
	// 		taskResponse := responses.NewTaskResponse(task)
	// 		taskResponses = append(taskResponses, taskResponse)
	// 	}
	// 	return taskResponses, nil
	// }
	return nil, nil

}
