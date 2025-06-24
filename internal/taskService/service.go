package taskService

import (
	"Tasks/internal/models"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task models.Task) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id string) (models.Task, error)
	UpdateTask(task models.Task) (models.Task, error)
	DeleteTask(id string) error
	GetTasksByUserID(userID string) ([]models.Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) GetTasksByUserID(userID string) ([]models.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *taskService) CreateTask(task models.Task) (models.Task, error) {
	task.ID = uuid.NewString()

	if err := s.repo.CreateTask(task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskById(id string) (models.Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTask(task models.Task) (models.Task, error) {
	if err := s.repo.UpdateTask(task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
