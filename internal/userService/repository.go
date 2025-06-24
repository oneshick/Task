package userService

import (
	"Tasks/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u models.User) error
	GetAllUser() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	UpdateUser(u models.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(u models.User) error {
	return r.db.Create(&u).Error
}

func (r *userRepository) GetAllUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateUser(u models.User) error {
	return r.db.Save(&u).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
