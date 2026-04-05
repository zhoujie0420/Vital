package dao

import (
	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/utils"

	"gorm.io/gorm"
)

type DietDAO struct{ db *gorm.DB }

func NewDietDAO() *DietDAO { return &DietDAO{db: utils.GetDB()} }

func (d *DietDAO) CreateRecord(r *model.DietRecord) error { return d.db.Create(r).Error }

func (d *DietDAO) FindByUserID(userID uint, page, pageSize int) ([]model.DietRecord, int64, error) {
	var records []model.DietRecord
	var total int64
	d.db.Model(&model.DietRecord{}).Where("user_id = ?", userID).Count(&total)
	err := d.db.Where("user_id = ?", userID).Order("record_date DESC, created_at DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	return records, total, err
}

// Food 相关
func (d *DietDAO) FindFoods(keyword string, category string, userID uint) ([]model.Food, error) {
	var foods []model.Food
	query := d.db.Where("user_id = 0 OR user_id = ?", userID)
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	if category != "" && category != "全部" {
		query = query.Where("category = ?", category)
	}
	err := query.Order("is_custom ASC, category, name").Limit(100).Find(&foods).Error
	return foods, err
}

func (d *DietDAO) CreateFood(f *model.Food) error { return d.db.Create(f).Error }
