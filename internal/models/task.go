package models

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID string `json:"user_id"` // Добавляем явное поле UserID

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskUpdateRequest struct {
	Title  *string `json:"title,omitempty"`
	Status *string `json:"status,omitempty"`
}
