package model

import (
	"time"

	"gorm.io/gorm"
)

// WeightRecord 体重记录模型
type WeightRecord struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	UserID      uint      `gorm:"not null;index" json:"user_id"`
	RecordDate  time.Time `gorm:"not null;index" json:"record_date"`
	Weight      float64   `gorm:"type:decimal(5,2);not null" json:"weight"` // 体重(kg)
	Height      float64   `gorm:"type:decimal(5,2)" json:"height,omitempty"` // 身高(cm)
	BMI         float64   `gorm:"type:decimal(5,2);index" json:"bmi,omitempty"` // BMI值
}

// TableName 设置表名
func (WeightRecord) TableName() string {
	return "weight_records"
}

// CreateWeightRequest 创建体重记录请求
type CreateWeightRequest struct {
	RecordDate time.Time `json:"record_date" binding:"required"`
	Weight     float64   `json:"weight" binding:"required,gt=0"`
	Height     float64   `json:"height,omitempty" binding:"omitempty,gt=0"`
}

// WeightResponse 体重记录响应
type WeightResponse struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	RecordDate time.Time `json:"record_date"`
	Weight     float64   `json:"weight"`
	Height     float64   `json:"height,omitempty"`
	BMI        float64   `json:"bmi,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// CalculateBMI 计算BMI值
func (w *WeightRecord) CalculateBMI() float64 {
	if w.Height > 0 {
		heightInMeter := w.Height / 100 // 转换为米
		w.BMI = w.Weight / (heightInMeter * heightInMeter)
		return w.BMI
	}
	return 0
}