package taskService

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}
