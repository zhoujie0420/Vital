package model

import (
	"time"

	"gorm.io/gorm"
)

// DietRecord 饮食记录模型
type DietRecord struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	UserID        uint      `gorm:"not null;index" json:"user_id"`
	RecordDate    time.Time `gorm:"not null;index" json:"record_date"`
	MealType      string    `gorm:"size:20;not null;index" json:"meal_type"`     // 餐次类型: breakfast,lunch,dinner,snack
	FoodItems     string    `gorm:"type:text" json:"food_items"`                 // 食物列表(JSON格式)
	TotalCalories float64   `gorm:"type:decimal(10,2)" json:"total_calories"`    // 总卡路里
	Protein       float64   `gorm:"type:decimal(10,2)" json:"protein,omitempty"` // 蛋白质(g)
	Carbs         float64   `gorm:"type:decimal(10,2)" json:"carbs,omitempty"`   // 碳水化合物(g)
	Fat           float64   `gorm:"type:decimal(10,2)" json:"fat,omitempty"`     // 脂肪(g)
	WaterIntake   int       `json:"water_intake,omitempty"`                      // 饮水量(ml)
	Notes         string    `gorm:"type:text" json:"notes,omitempty"`            // 备注
}

// TableName 设置表名
func (DietRecord) TableName() string {
	return "diet_records"
}

// CreateDietRequest 创建饮食记录请求
type CreateDietRequest struct {
	RecordDate    time.Time `json:"record_date" binding:"required"`
	MealType      string    `json:"meal_type" binding:"required,oneof=breakfast lunch dinner snack"`
	FoodItems     string    `json:"food_items" binding:"required"` // JSON格式的食物列表
	TotalCalories float64   `json:"total_calories" binding:"required"`
	Protein       float64   `json:"protein,omitempty"`
	Carbs         float64   `json:"carbs,omitempty"`
	Fat           float64   `json:"fat,omitempty"`
	WaterIntake   int       `json:"water_intake,omitempty"`
	Notes         string    `json:"notes,omitempty"`
}

// DietResponse 饮食记录响应
type DietResponse struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"user_id"`
	RecordDate    time.Time `json:"record_date"`
	MealType      string    `json:"meal_type"`
	FoodItems     string    `json:"food_items"`
	TotalCalories float64   `json:"total_calories"`
	Protein       float64   `json:"protein,omitempty"`
	Carbs         float64   `json:"carbs,omitempty"`
	Fat           float64   `json:"fat,omitempty"`
	WaterIntake   int       `json:"water_intake,omitempty"`
	Notes         string    `json:"notes,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

// MealType 餐次类型的常量定义
const (
	MealBreakfast = "breakfast" // 早餐
	MealLunch     = "lunch"     // 午餐
	MealDinner    = "dinner"    // 晚餐
	MealSnack     = "snack"     // 加餐/零食
)

// UpdateDietRequest 更新饮食记录请求
type UpdateDietRequest struct {
	MealType      string  `json:"meal_type,omitempty"`
	FoodItems     string  `json:"food_items,omitempty"`
	TotalCalories float64 `json:"total_calories,omitempty"`
	Protein       float64 `json:"protein,omitempty"`
	Carbs         float64 `json:"carbs,omitempty"`
	Fat           float64 `json:"fat,omitempty"`
	WaterIntake   int     `json:"water_intake,omitempty"`
	Notes         string  `json:"notes,omitempty"`
}
