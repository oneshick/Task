package models

import "time"

type User struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	Tasks []Task `gorm:"foreignKey:UserID"`
}
type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" `
}

type UserUpdateRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}
