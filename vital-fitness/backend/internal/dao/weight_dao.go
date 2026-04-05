package dao

import (
	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/utils"

	"gorm.io/gorm"
)

type WeightDAO struct{ db *gorm.DB }

func NewWeightDAO() *WeightDAO { return &WeightDAO{db: utils.GetDB()} }

func (d *WeightDAO) Create(r *model.WeightRecord) error { return d.db.Create(r).Error }

func (d *WeightDAO) FindByUserID(userID uint, page, pageSize int) ([]model.WeightRecord, int64, error) {
	var records []model.WeightRecord
	var total int64
	d.db.Model(&model.WeightRecord{}).Where("user_id = ?", userID).Count(&total)
	err := d.db.Where("user_id = ?", userID).Order("record_date DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	return records, total, err
}

func (d *WeightDAO) FindLatest(userID uint) (*model.WeightRecord, error) {
	var r model.WeightRecord
	err := d.db.Where("user_id = ?", userID).Order("record_date DESC").First(&r).Error
	return &r, err
}
