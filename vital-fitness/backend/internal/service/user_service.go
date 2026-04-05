package service

import (
	"errors"

	"vital-fitness/backend/internal/dao"
	"vital-fitness/backend/internal/model"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService() *UserService {
	return &UserService{userDAO: dao.NewUserDAO()}
}

func (s *UserService) GetProfile(userID uint) (*model.UserProfile, error) {
	user, err := s.userDAO.FindByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return &model.UserProfile{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Gender:    user.Gender,
		Height:    user.Height,
		Weight:    user.Weight,
		CreatedAt: user.CreatedAt,
	}, nil
}
