package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Username string `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email    string `gorm:"size:100;uniqueIndex" json:"email"`
	Phone    string `gorm:"size:20;uniqueIndex" json:"phone,omitempty"`
	Password string `gorm:"size:255;not null" json:"-"`
	Avatar   string `gorm:"size:255" json:"avatar,omitempty"`
	Nickname string `gorm:"size:50" json:"nickname,omitempty"`
	Gender   int    `gorm:"default:0" json:"gender,omitempty"` // 0:未知 1:男 2:女
	Birthday time.Time `json:"birthday,omitempty"`
	Height   float64   `json:"height,omitempty"` // 身高(cm)
	Weight   float64   `json:"weight,omitempty"` // 体重(kg)
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
}

// UserProfile 用户资料返回结构
type UserProfile struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Gender    int       `json:"gender,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	Height    float64   `json:"height,omitempty"`
	Weight    float64   `json:"weight,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token     string      `json:"token"`
	User      UserProfile `json:"user"`
	ExpiresAt time.Time   `json:"expires_at"`
}