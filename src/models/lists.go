package models

type List struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Color string `json:"color"`
}
