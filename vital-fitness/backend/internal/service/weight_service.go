package service

import (
	"errors"

	"vital-fitness/backend/internal/dao"
	"vital-fitness/backend/internal/model"
)

type WeightService struct{ weightDAO *dao.WeightDAO }

func NewWeightService() *WeightService { return &WeightService{weightDAO: dao.NewWeightDAO()} }

func (s *WeightService) Create(userID uint, req *model.CreateWeightRequest) (*model.WeightRecord, error) {
	r := &model.WeightRecord{
		UserID:     userID,
		RecordDate: req.RecordDate,
		Weight:     req.Weight,
		Height:     req.Height,
	}
	r.CalculateBMI()
	if err := s.weightDAO.Create(r); err != nil {
		return nil, errors.New("创建体重记录失败")
	}
	return r, nil
}

func (s *WeightService) GetList(userID uint, page, pageSize int) ([]model.WeightRecord, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	return s.weightDAO.FindByUserID(userID, page, pageSize)
}
