package userService

import "github.com/google/uuid"

type UserService interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(user User) (User, error) {
	user.ID = uuid.NewString()

	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUser()
}

func (s *userService) GetUserByID(id string) (User, error) {
	return s.repo.GetUserById(id)
}

func (s *userService) UpdateUser(user User) (User, error) {
	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
