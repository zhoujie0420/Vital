package dao

import (
	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/utils"

	"gorm.io/gorm"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 创建 UserDAO 实例
func NewUserDAO() *UserDAO {
	return &UserDAO{db: utils.GetDB()}
}

// Create 创建用户
func (d *UserDAO) Create(user *model.User) error {
	return d.db.Create(user).Error
}

// FindByUsername 根据用户名查找用户
func (d *UserDAO) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (d *UserDAO) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := d.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID 根据 ID 查找用户
func (d *UserDAO) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := d.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ExistsByUsername 检查用户名是否已存在
func (d *UserDAO) ExistsByUsername(username string) bool {
	var count int64
	d.db.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

// ExistsByEmail 检查邮箱是否已存在
func (d *UserDAO) ExistsByEmail(email string) bool {
	var count int64
	d.db.Model(&model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// FindByOpenID 根据微信 openid 查找用户
func (d *UserDAO) FindByOpenID(openID string) (*model.User, error) {
	var user model.User
	err := d.db.Where("open_id = ?", openID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
