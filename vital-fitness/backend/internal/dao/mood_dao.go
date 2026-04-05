package dao

import (
	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/utils"

	"gorm.io/gorm"
)

type MoodDAO struct{ db *gorm.DB }

func NewMoodDAO() *MoodDAO { return &MoodDAO{db: utils.GetDB()} }

func (d *MoodDAO) Create(r *model.MoodRecord) error { return d.db.Create(r).Error }

func (d *MoodDAO) FindByUserID(userID uint, page, pageSize int) ([]model.MoodRecord, int64, error) {
	var records []model.MoodRecord
	var total int64
	d.db.Model(&model.MoodRecord{}).Where("user_id = ?", userID).Count(&total)
	err := d.db.Where("user_id = ?", userID).Order("record_date DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).Find(&records).Error
	return records, total, err
}
