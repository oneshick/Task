package taskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(title, status string) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTask(id, title, status string) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(title, status string) (Task, error) {
	task := Task{
		ID:     uuid.NewString(),
		Title:  title,
		Status: status,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskById(id string) (Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTask(id, title, status string) (Task, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	task.Title = title
	task.Status = status

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
