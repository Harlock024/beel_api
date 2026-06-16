package repositories

import (
	"beel_api/src/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTasksByListId(list_id uuid.UUID) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.Preload("Tags").Preload("Subtasks.Tags").Where("list_id = ? AND parent_id IS NULL", list_id).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
func (r *TaskRepository) CreateTask(task *models.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}
func (r *TaskRepository) UpdateTask(task *models.Task) error {
	if err := r.db.Save(task).Error; err != nil {
		return err
	}
	return nil
}
func (r *TaskRepository) DeleteTask(task *models.Task) error {
	if err := r.db.Delete(task).Error; err != nil {
		return err
	}
	return nil
}
func (r *TaskRepository) GetTaskById(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	if err := r.db.Preload("Tags").Preload("Subtasks.Tags").Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) GetTasks(user_id uuid.UUID) ([]*models.Task,error){
	var tasks []*models.Task
	if err := r.db.Preload("Tags").Preload("Subtasks.Tags").Where("user_id = ?",user_id).Find(&tasks).Error; err != nil {
		return  nil,err
	}
	return tasks ,nil
} 

func (r *TaskRepository) GetTasksByFilter(start string, end string, user_id uuid.UUID) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.
		Preload("Tags").
		Preload("Subtasks.Tags").
		Where("user_id = ? AND due_date >= ? AND due_date < ? AND parent_id IS NULL", user_id, start, end).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetSubtasksByParentId(parentId uuid.UUID) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.Preload("Tags").Where("parent_id = ?", parentId).Order("created_at ASC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByIdRaw(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	if err := r.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) IsAncestor(taskId uuid.UUID, potentialAncestorId uuid.UUID) (bool, error) {
	currentId := taskId
	for {
		var task models.Task
		if err := r.db.Select("parent_id").Where("id = ?", currentId).First(&task).Error; err != nil {
			return false, err
		}
		if task.ParentID == nil {
			return false, nil
		}
		if *task.ParentID == potentialAncestorId {
			return true, nil
		}
		currentId = *task.ParentID
	}
}

func (r *TaskRepository) GetTasksByTag(tagId uuid.UUID) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.
		Joins("JOIN task_tags ON task_tags.task_id = tasks.id").
		Where("task_tags.tag_id = ? AND tasks.parent_id IS NULL", tagId).
		Preload("Tags").
		Preload("Subtasks.Tags").
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) AddTagToTask(taskId uuid.UUID, tagId uuid.UUID) error {
	task := &models.Task{ID: taskId}
	tag := &models.Tag{ID: tagId}
	return r.db.Model(task).Association("Tags").Append(tag)
}

func (r *TaskRepository) RemoveTagFromTask(taskId uuid.UUID, tagId uuid.UUID) error {
	task := &models.Task{ID: taskId}
	tag := &models.Tag{ID: tagId}
	return r.db.Model(task).Association("Tags").Delete(tag)
}

func (r *TaskRepository) GetTaskByIdWithTags(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	if err := r.db.Preload("Tags").Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) BatchUpdateTasks(tasks []models.Task) error {
	return r.db.Save(&tasks).Error
}

func (r *TaskRepository) CountTasksByUserId(userId uuid.UUID) (int64, error) {
	var count int64
	if err := r.db.Model(&models.Task{}).Where("user_id = ? AND parent_id IS NULL", userId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
