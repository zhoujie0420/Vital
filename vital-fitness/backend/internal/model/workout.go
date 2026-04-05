package model

import (
	"time"

	"gorm.io/gorm"
)

// Exercise 动作模型
type Exercise struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Name        string `gorm:"size:100;not null;index" json:"name"`
	Category    string `gorm:"size:50;index" json:"category"` // 动作分类: 胸、背、腿、肩、手臂、核心等
	Description string `gorm:"type:text" json:"description,omitempty"`
	VideoURL    string `gorm:"size:255" json:"video_url,omitempty"`
	ImageURL    string `gorm:"size:255" json:"image_url,omitempty"`
}

// Workout 训练记录模型
type Workout struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	UserID      uint      `gorm:"not null;index" json:"user_id"`
	ExerciseID  uint      `gorm:"not null;index" json:"exercise_id"`
	Exercise    Exercise  `gorm:"foreignKey:ExerciseID" json:"exercise"`
	WorkoutDate time.Time `gorm:"not null;index" json:"workout_date"`

	Weight   float64 `gorm:"type:decimal(10,2)" json:"weight"` // 重量(kg)
	Sets     int     `gorm:"not null;default:1" json:"sets"`   // 组数
	Reps     int     `gorm:"not null;default:1" json:"reps"`   // 次数
	RestTime int     `json:"rest_time,omitempty"`              // 休息时间(秒)

	Feeling  int    `gorm:"default:3" json:"feeling"`            // 训练感受(1-5分)
	Notes    string `gorm:"type:text" json:"notes,omitempty"`    // 备注
	ImageURL string `gorm:"size:255" json:"image_url,omitempty"` // 训练照片
}

// TableName 设置表名
func (Exercise) TableName() string {
	return "exercises"
}

// TableName 设置表名
func (Workout) TableName() string {
	return "workouts"
}

// CreateWorkoutRequest 创建训练记录请求
type CreateWorkoutRequest struct {
	ExerciseID  uint      `json:"exercise_id" binding:"required"`
	WorkoutDate time.Time `json:"workout_date" binding:"required"`
	Weight      float64   `json:"weight" binding:"required"`
	Sets        int       `json:"sets" binding:"required,min=1"`
	Reps        int       `json:"reps" binding:"required,min=1"`
	RestTime    int       `json:"rest_time,omitempty"`
	Feeling     int       `json:"feeling,omitempty" binding:"min=1,max=5"`
	Notes       string    `json:"notes,omitempty"`
}

// UpdateWorkoutRequest 更新训练记录请求
type UpdateWorkoutRequest struct {
	ExerciseID  uint      `json:"exercise_id,omitempty"`
	WorkoutDate time.Time `json:"workout_date,omitempty"`
	Weight      float64   `json:"weight,omitempty"`
	Sets        int       `json:"sets,omitempty" binding:"min=1"`
	Reps        int       `json:"reps,omitempty" binding:"min=1"`
	RestTime    int       `json:"rest_time,omitempty"`
	Feeling     int       `json:"feeling,omitempty" binding:"min=1,max=5"`
	Notes       string    `json:"notes,omitempty"`
}

// WorkoutResponse 训练记录响应
type WorkoutResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Exercise    Exercise  `json:"exercise"`
	WorkoutDate time.Time `json:"workout_date"`
	Weight      float64   `json:"weight"`
	Sets        int       `json:"sets"`
	Reps        int       `json:"reps"`
	RestTime    int       `json:"rest_time,omitempty"`
	Feeling     int       `json:"feeling"`
	Notes       string    `json:"notes,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
