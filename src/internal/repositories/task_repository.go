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
	if err := r.db.Preload("Tags").Where("list_id = ?", list_id).Find(&tasks).Error; err != nil {
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
	if err := r.db.Preload("Tags").Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
