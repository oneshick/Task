package taskService

import (
	"Tasks/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(t models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id string) (models.Task, error)
	UpdateTask(t models.Task) error
	DeleteTask(id string) error
	GetTasksByUserID(userID string) ([]models.Task, error)
}

type tRepository struct {
	db *gorm.DB
}

func NewTRepository(db *gorm.DB) TaskRepository {
	return &tRepository{db: db}
}
func (r *tRepository) CreateTask(t models.Task) error {
	return r.db.Create(&t).Error
}

func (r *tRepository) GetTasksByUserID(userID string) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *tRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *tRepository) GetTaskById(id string) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *tRepository) UpdateTask(t models.Task) error {
	return r.db.Save(&t).Error
}

func (r *tRepository) DeleteTask(id string) error {
	return r.db.Delete(&models.Task{}, "id = ?", id).Error
}
