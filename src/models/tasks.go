package models

type Task struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
	ListID      string `json:"listId"`
}
