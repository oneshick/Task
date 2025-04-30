package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(t Task) error
	GetAllTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTask(t Task) error
	DeleteTask(id string) error
}

type tRepository struct {
	db *gorm.DB
}

func NewTRepository(db *gorm.DB) TaskRepository {
	return &tRepository{db: db}
}
func (r *tRepository) CreateTask(t Task) error {
	return r.db.Create(&t).Error
}
func (r *tRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *tRepository) GetTaskById(id string) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *tRepository) UpdateTask(t Task) error {
	return r.db.Save(&t).Error
}

func (r *tRepository) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
