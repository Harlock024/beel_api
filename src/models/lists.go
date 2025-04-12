package models

type List struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Color  string `json:"color"`
	UserID string `json:"user_id" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	User   User   `gorm:"foreignKey:UserID"`
	Tasks  []Task
}
