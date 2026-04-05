package service

import (
	"errors"

	"vital-fitness/backend/internal/dao"
	"vital-fitness/backend/internal/model"
)

type MoodService struct{ moodDAO *dao.MoodDAO }

func NewMoodService() *MoodService { return &MoodService{moodDAO: dao.NewMoodDAO()} }

func (s *MoodService) Create(userID uint, req *model.CreateMoodRequest) (*model.MoodRecord, error) {
	r := &model.MoodRecord{
		UserID:      userID,
		RecordDate:  req.RecordDate,
		MoodScore:   req.MoodScore,
		MoodTags:    req.MoodTags,
		Description: req.Description,
	}
	if err := s.moodDAO.Create(r); err != nil {
		return nil, errors.New("创建心情记录失败")
	}
	return r, nil
}

func (s *MoodService) GetList(userID uint, page, pageSize int) ([]model.MoodRecord, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	return s.moodDAO.FindByUserID(userID, page, pageSize)
}
