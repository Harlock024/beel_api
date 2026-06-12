package services

import (
	"beel_api/src/api/responses"
	"beel_api/src/dtos"
	"beel_api/src/internal/models"
	"beel_api/src/internal/repositories"
	"fmt"
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
		ID:       uuid.New(),
		Title:    task.Title,
		ListID:   list_id,
		UserID:   user_id,
		ParentID: task.ParentID,
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

	if task.ParentID != nil {
		if *task.ParentID == task_id {
			return nil, fmt.Errorf("a task cannot be its own parent")
		}

		parent, err := s.repo.GetTaskByIdRaw(*task.ParentID)
		if err != nil {
			return nil, fmt.Errorf("parent task not found")
		}
		if parent.ListID != existingTask.ListID {
			return nil, fmt.Errorf("parent task must be in the same list")
		}

		isAncestor, err := s.repo.IsAncestor(*task.ParentID, task_id)
		if err != nil {
			return nil, err
		}
		if isAncestor {
			return nil, fmt.Errorf("cannot set parent: would create a cycle")
		}

		existingTask.ParentID = task.ParentID
	} else if task.ParentID != nil && *task.ParentID == uuid.Nil {
		existingTask.ParentID = nil
	}

	if task.ColumnID != nil {
		existingTask.ColumnID = task.ColumnID
	}
	if task.Position != nil {
		existingTask.Position = *task.Position
	}

	err = s.repo.UpdateTask(existingTask)
	if err != nil {
		return nil, err
	}

	updatedTask, err := s.repo.GetTaskById(task_id)
	if err != nil {
		return responses.NewTaskResponse(existingTask), nil
	}

	return responses.NewTaskResponse(updatedTask), nil
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

func (s TaskService) GetTasksByFilter(filter string, user_id uuid.UUID) ([]*responses.TaskResponse, error) {
	if filter == "" {
		return nil, nil
	}

	loc, _ := time.LoadLocation("America/Mexico_City")

	if filter == "today" {
		now := time.Now().In(loc)
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).UTC()
		end := start.Add(24 * time.Hour).UTC()
		tasks, err := s.repo.GetTasksByFilter(start.String(), end.String(), user_id)
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
		tasks, err := s.repo.GetTasksByFilter(start.String(), end.String(), user_id)
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

func (s *TaskService) GetSubtasksByParentId(parentId uuid.UUID) ([]*responses.TaskResponse, error) {
	subtasks, err := s.repo.GetSubtasksByParentId(parentId)
	if err != nil {
		return nil, err
	}
	var taskResponses []*responses.TaskResponse
	for _, subtask := range subtasks {
		taskResponses = append(taskResponses, responses.NewTaskResponse(subtask))
	}
	if len(taskResponses) == 0 {
		return []*responses.TaskResponse{}, nil
	}
	return taskResponses, nil
}

func (s *TaskService) AddSubtask(parentId uuid.UUID, dto *dtos.NewTaskDTO, user_id uuid.UUID) (*responses.TaskResponse, error) {
	parent, err := s.repo.GetTaskByIdRaw(parentId)
	if err != nil {
		return nil, err
	}

	newSubtask := &models.Task{
		ID:       uuid.New(),
		Title:    dto.Title,
		ListID:   parent.ListID,
		UserID:   user_id,
		ParentID: &parentId,
	}
	if err := s.repo.CreateTask(newSubtask); err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(newSubtask), nil
}

func (s *TaskService) RemoveSubtask(parentId uuid.UUID, subtaskId uuid.UUID) error {
	subtask, err := s.repo.GetTaskByIdRaw(subtaskId)
	if err != nil {
		return err
	}
	if subtask.ParentID == nil || *subtask.ParentID != parentId {
		return err
	}
	return s.repo.DeleteTask(subtask)
}

func (s *TaskService) GetTasksByTag(tagId uuid.UUID) ([]*responses.TaskResponse, error) {
	tasks, err := s.repo.GetTasksByTag(tagId)
	if err != nil {
		return nil, err
	}
	var taskResponses []*responses.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, responses.NewTaskResponse(task))
	}
	if len(taskResponses) == 0 {
		return []*responses.TaskResponse{}, nil
	}
	return taskResponses, nil
}

func (s *TaskService) AddTagToTask(taskId uuid.UUID, tagId uuid.UUID) (*responses.TaskResponse, error) {
	if err := s.repo.AddTagToTask(taskId, tagId); err != nil {
		return nil, err
	}
	task, err := s.repo.GetTaskByIdWithTags(taskId)
	if err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(task), nil
}

func (s *TaskService) RemoveTagFromTask(taskId uuid.UUID, tagId uuid.UUID) (*responses.TaskResponse, error) {
	if err := s.repo.RemoveTagFromTask(taskId, tagId); err != nil {
		return nil, err
	}
	task, err := s.repo.GetTaskByIdWithTags(taskId)
	if err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(task), nil
}

func (s *TaskService) BatchUpdateTasks(dto *dtos.BatchUpdateDTO) error {
	var tasks []models.Task
	for _, item := range dto.Tasks {
		task := models.Task{
			ID:       item.ID,
			ColumnID: item.ColumnID,
		}
		if item.Position != nil {
			task.Position = *item.Position
		}
		tasks = append(tasks, task)
	}
	return s.repo.BatchUpdateTasks(tasks)
}
