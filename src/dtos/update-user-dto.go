package dtos


type UpdateUser struct {
	Username  string `json:"username" binding:"required"`
	AvatarUrl string `json:"avatar_url" binding:"required"`
}
