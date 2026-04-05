package model

import (
"time"

"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	OpenID    string         `gorm:"size:100;uniqueIndex;column:open_id" json:"-"`
	Username  string         `gorm:"size:50;uniqueIndex" json:"username"`
	Email     string         `gorm:"size:100;index" json:"email"`
	Phone     string         `gorm:"size:20;index" json:"phone,omitempty"`
	Password  string         `gorm:"size:255" json:"-"`
	Avatar    string         `gorm:"size:255" json:"avatar,omitempty"`
	Nickname  string         `gorm:"size:50" json:"nickname,omitempty"`
	Gender    int            `gorm:"default:0" json:"gender,omitempty"`
	Birthday  *time.Time     `json:"birthday,omitempty"`
	Height    float64        `json:"height,omitempty"`
	Weight    float64        `json:"weight,omitempty"`
}

func (User) TableName() string { return "users" }

type UserProfile struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Avatar    string     `json:"avatar,omitempty"`
	Nickname  string     `json:"nickname,omitempty"`
	Gender    int        `json:"gender,omitempty"`
	Height    float64    `json:"height,omitempty"`
	Weight    float64    `json:"weight,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

type WxLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

type LoginResponse struct {
	Token     string      `json:"token"`
	User      UserProfile `json:"user"`
	ExpiresAt time.Time   `json:"expires_at"`
}
