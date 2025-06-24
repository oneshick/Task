package userService

import (
	"Tasks/internal/models"
	"Tasks/internal/taskService"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) error
	GetTasksForUser(userID string) ([]models.Task, error)
}

type userService struct {
	repo     UserRepository
	taskRepo taskService.TaskRepository
}

func NewUserService(r UserRepository, tr taskService.TaskRepository) UserService {
	return &userService{repo: r, taskRepo: tr}
}

func (s *userService) GetTasksForUser(userID string) ([]models.Task, error) {
	return s.taskRepo.GetTasksByUserID(userID)
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	user.ID = uuid.NewString()

	if err := s.repo.CreateUser(user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUser()
}

func (s *userService) GetUserByID(id string) (models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *userService) UpdateUser(user models.User) (models.User, error) {
	if err := s.repo.UpdateUser(user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
