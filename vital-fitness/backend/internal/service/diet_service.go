package service

import (
	"errors"

	"vital-fitness/backend/internal/dao"
	"vital-fitness/backend/internal/model"
)

type DietService struct{ dietDAO *dao.DietDAO }

func NewDietService() *DietService { return &DietService{dietDAO: dao.NewDietDAO()} }

func (s *DietService) CreateRecord(userID uint, req *model.CreateDietRequest) (*model.DietRecord, error) {
	r := &model.DietRecord{
		UserID:        userID,
		RecordDate:    req.RecordDate,
		MealType:      req.MealType,
		FoodItems:     req.FoodItems,
		TotalCalories: req.TotalCalories,
		Protein:       req.Protein,
		Carbs:         req.Carbs,
		Fat:           req.Fat,
		Notes:         req.Notes,
	}
	if err := s.dietDAO.CreateRecord(r); err != nil {
		return nil, errors.New("创建饮食记录失败")
	}
	return r, nil
}

func (s *DietService) DeleteRecord(id uint, userID uint) error {
	return s.dietDAO.Delete(id, userID)
}

func (s *DietService) GetRecords(userID uint, page, pageSize int) ([]model.DietRecord, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	return s.dietDAO.FindByUserID(userID, page, pageSize)
}

func (s *DietService) GetFoods(keyword, category string, userID uint) ([]model.Food, error) {
	return s.dietDAO.FindFoods(keyword, category, userID)
}

func (s *DietService) CreateFood(userID uint, req *model.CreateFoodRequest) (*model.Food, error) {
	f := &model.Food{
		Name:     req.Name,
		Category: req.Category,
		Calories: req.Calories,
		Protein:  req.Protein,
		Carbs:    req.Carbs,
		Fat:      req.Fat,
		Unit:     req.Unit,
		IsCustom: true,
		UserID:   userID,
	}
	if f.Unit == "" {
		f.Unit = "100g"
	}
	if err := s.dietDAO.CreateFood(f); err != nil {
		return nil, errors.New("创建食物失败")
	}
	return f, nil
}
