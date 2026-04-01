package service

import (
	"errors"

	"vital-fitness/backend/internal/dao"
	"vital-fitness/backend/internal/model"

	"golang.org/x/crypto/bcrypt"
)

// UserService 用户业务逻辑
type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 创建 UserService 实例
func NewUserService() *UserService {
	return &UserService{userDAO: dao.NewUserDAO()}
}

// Register 用户注册
func (s *UserService) Register(req *model.RegisterRequest) (*model.UserProfile, error) {
	// 检查用户名是否已存在
	if s.userDAO.ExistsByUsername(req.Username) {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	if s.userDAO.ExistsByEmail(req.Email) {
		return nil, errors.New("邮箱已被注册")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Nickname: req.Username,
	}

	if err := s.userDAO.Create(user); err != nil {
		return nil, errors.New("注册失败，请稍后重试")
	}

	profile := &model.UserProfile{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt,
	}
	return profile, nil
}

// Login 用户登录
func (s *UserService) Login(req *model.LoginRequest) (*model.User, error) {
	// 支持用户名或邮箱登录
	user, err := s.userDAO.FindByUsername(req.Username)
	if err != nil {
		// 尝试用邮箱查找
		user, err = s.userDAO.FindByEmail(req.Username)
		if err != nil {
			return nil, errors.New("用户名或密码错误")
		}
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	return user, nil
}

// GetProfile 获取用户资料
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
		Birthday:  user.Birthday,
		Height:    user.Height,
		Weight:    user.Weight,
		CreatedAt: user.CreatedAt,
	}, nil
}
