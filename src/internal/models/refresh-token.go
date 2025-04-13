package models

type RefreshToken struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserID    uint   `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	IsRevoked bool   `json:"is_revoked"`
}
