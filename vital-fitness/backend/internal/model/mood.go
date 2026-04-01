package model

import (
	"time"

	"gorm.io/gorm"
)

// MoodRecord 心情记录模型
type MoodRecord struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	UserID      uint      `gorm:"not null;index" json:"user_id"`
	RecordDate  time.Time `gorm:"not null;index" json:"record_date"`
	MoodScore   int       `gorm:"not null;check:mood_score >= 1 AND mood_score <= 10" json:"mood_score"` // 心情分数(1-10)
	MoodTags    string    `gorm:"size:255" json:"mood_tags,omitempty"`                                   // 心情标签(逗号分隔)
	Description string    `gorm:"type:text" json:"description,omitempty"`                                // 描述
}

// TableName 设置表名
func (MoodRecord) TableName() string {
	return "mood_records"
}

// CreateMoodRequest 创建心情记录请求
type CreateMoodRequest struct {
	RecordDate  time.Time `json:"record_date" binding:"required"`
	MoodScore   int       `json:"mood_score" binding:"required,min=1,max=10"`
	MoodTags    string    `json:"mood_tags,omitempty"`
	Description string    `json:"description,omitempty"`
}

// MoodResponse 心情记录响应
type MoodResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	RecordDate  time.Time `json:"record_date"`
	MoodScore   int       `json:"mood_score"`
	MoodTags    string    `json:"mood_tags,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}