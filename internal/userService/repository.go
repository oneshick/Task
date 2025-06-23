package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(u User) error
	GetAllUser() ([]User, error)
	GetUserById(id string) (User, error)
	UpdateUser(u User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(u User) error {
	return r.db.Create(&u).Error
}

func (r *userRepository) GetAllUser() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id string) (User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateUser(u User) error {
	return r.db.Save(&u).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
