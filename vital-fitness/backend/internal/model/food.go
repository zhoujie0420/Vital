package model

import (
	"time"

	"gorm.io/gorm"
)

// Food 食物库
type Food struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Name     string  `gorm:"size:100;not null" json:"name"`
	Category string  `gorm:"size:50;index" json:"category"`      // 主食、肉类、蔬菜、水果、奶制品、零食、饮品
	Calories float64 `gorm:"type:decimal(10,2)" json:"calories"` // 每100g热量(kcal)
	Protein  float64 `gorm:"type:decimal(10,2)" json:"protein"`  // 每100g蛋白质(g)
	Carbs    float64 `gorm:"type:decimal(10,2)" json:"carbs"`    // 每100g碳水(g)
	Fat      float64 `gorm:"type:decimal(10,2)" json:"fat"`      // 每100g脂肪(g)
	Unit     string  `gorm:"size:20;default:'100g'" json:"unit"` // 计量单位
	IsCustom bool    `gorm:"default:false" json:"is_custom"`     // 是否用户自定义
	UserID   uint    `gorm:"index;default:0" json:"user_id"`     // 0=系统预置
}

func (Food) TableName() string { return "foods" }

// CreateFoodRequest 创建食物请求
type CreateFoodRequest struct {
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category"`
	Calories float64 `json:"calories" binding:"required"`
	Protein  float64 `json:"protein"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Unit     string  `json:"unit"`
}
