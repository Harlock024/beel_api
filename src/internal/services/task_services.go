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

func formatDueDate(dueDate string) string {
	if dueDate == "" {
		return ""
	}
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"02/01/2006",
		"01/02/2006",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05",
		time.RFC3339,
	}
	for _, layout := range formats {
		if t, err := time.Parse(layout, dueDate); err == nil {
			return t.Format("2006-01-02")
		}
	}
	return dueDate
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

func (s *TaskService) CreateTask(task *dtos.NewTaskDTO, list_id *uuid.UUID, user_id uuid.UUID) (*responses.TaskResponse, error) {

	newTask := &models.Task{
		ID:       uuid.New(),
		Title:    task.Title,
		ListID:   list_id,
		UserID:   user_id,
		DueDate:  formatDueDate(task.DueDate),
		ParentID: task.ParentID,
	}
	err := s.repo.CreateTask(newTask)
	if err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(newTask), nil
}

func (s *TaskService) UpdateTask(task_id uuid.UUID, userId uuid.UUID, task *dtos.UpdateTaskDTO) (*responses.TaskResponse, error) {

	existingTask, err := s.repo.GetTaskById(task_id)
	if err != nil {
		return nil, err
	}

	if existingTask.UserID != userId {
		return nil, fmt.Errorf("forbidden: you do not own this task")
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
		existingTask.DueDate = formatDueDate(task.DueDate)
	}
	if task.ListID != nil {
		existingTask.ListID = task.ListID
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
		if parent.ListID == nil || existingTask.ListID == nil || *parent.ListID != *existingTask.ListID {
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
func (s *TaskService) DeleteTask(task_id uuid.UUID, userId uuid.UUID) error {

	existingTask, err := s.repo.GetTaskById(task_id)
	if err != nil {
		return err
	}

	if existingTask.UserID != userId {
		return fmt.Errorf("forbidden: you do not own this task")
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

func (s TaskService) GetTasks(filter string, user_id uuid.UUID) ([]*responses.TaskResponse, error) {
	if filter != "" {
	loc, _ := time.LoadLocation("America/Mexico_City")
	
	switch filter {
	case  "today" :
		start :=time.Now()
		end :=start.AddDate(0,0,1)
		tasks, err := s.repo.GetTasksByFilter(start.Format(time.DateOnly),end.Format(time.DateOnly) , user_id)
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
	
	case "upcoming" :
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
		return taskResponses,nil
	}
		tasks,err := s.repo.GetTasks(user_id)
		if err != nil {
			return nil,err
		}
		if len (tasks) == 0 {
			return nil,nil
		}
		var taskResponses []*responses.TaskResponse
		for  _ ,task := range tasks {
			taskResponse := responses.NewTaskResponse(task)
			taskResponses = append(taskResponses,taskResponse)
		}

	}
		

	return nil,nil
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

func (s *TaskService) AddSubtask(parentId uuid.UUID, userId uuid.UUID, dto *dtos.NewTaskDTO, subtaskUserId uuid.UUID) (*responses.TaskResponse, error) {
	parent, err := s.repo.GetTaskByIdRaw(parentId)
	if err != nil {
		return nil, err
	}

	if parent.UserID != userId {
		return nil, fmt.Errorf("forbidden: you do not own this task")
	}

	newSubtask := &models.Task{
		ID:       uuid.New(),
		Title:    dto.Title,
		ListID:   parent.ListID,
		UserID:   subtaskUserId,
		ParentID: &parentId,
	}
	if err := s.repo.CreateTask(newSubtask); err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(newSubtask), nil
}

func (s *TaskService) RemoveSubtask(parentId uuid.UUID, subtaskId uuid.UUID, userId uuid.UUID) error {
	parent, err := s.repo.GetTaskByIdRaw(parentId)
	if err != nil {
		return err
	}

	if parent.UserID != userId {
		return fmt.Errorf("forbidden: you do not own this task")
	}

	subtask, err := s.repo.GetTaskByIdRaw(subtaskId)
	if err != nil {
		return err
	}

	if subtask.UserID != userId {
		return fmt.Errorf("forbidden: you do not own this subtask")
	}

	if subtask.ParentID == nil || *subtask.ParentID != parentId {
		return fmt.Errorf("subtask does not belong to the specified parent")
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

func (s *TaskService) AddTagToTask(taskId uuid.UUID, tagId uuid.UUID, userId uuid.UUID) (*responses.TaskResponse, error) {
	task, err := s.repo.GetTaskByIdRaw(taskId)
	if err != nil {
		return nil, err
	}

	if task.UserID != userId {
		return nil, fmt.Errorf("forbidden: you do not own this task")
	}

	if err := s.repo.AddTagToTask(taskId, tagId); err != nil {
		return nil, err
	}
	task, err = s.repo.GetTaskByIdWithTags(taskId)
	if err != nil {
		return nil, err
	}
	return responses.NewTaskResponse(task), nil
}

func (s *TaskService) RemoveTagFromTask(taskId uuid.UUID, tagId uuid.UUID, userId uuid.UUID) (*responses.TaskResponse, error) {
	task, err := s.repo.GetTaskByIdRaw(taskId)
	if err != nil {
		return nil, err
	}

	if task.UserID != userId {
		return nil, fmt.Errorf("forbidden: you do not own this task")
	}

	if err := s.repo.RemoveTagFromTask(taskId, tagId); err != nil {
		return nil, err
	}
	task, err = s.repo.GetTaskByIdWithTags(taskId)
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
