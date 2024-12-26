package service

import (
	"errors"

	"github.com/lwmacct/241224-go-template-gin/app/api/model"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// NewUserService 构造函数
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// CheckLogin 根据用户名和密码校验用户
func (s *UserService) CheckLogin(username, password string) (*model.User, error) {
	var user model.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("password incorrect")
	}
	return &user, nil
}

// GetUserByID 根据 ID 获取用户信息
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建新用户
func (s *UserService) CreateUser(u *model.User) error {
	return s.DB.Create(u).Error
}
