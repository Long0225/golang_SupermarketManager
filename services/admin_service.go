package services

import (
	"errors"
	"github.com/supermarketmanager/database"
	"github.com/supermarketmanager/models"
	"gorm.io/gorm"
)

// AdminService 管理员服务

type AdminService struct {
	DB *gorm.DB
}

// NewAdminService 创建管理员服务实例
func NewAdminService() *AdminService {
	return &AdminService{
		DB: database.GetDB(),
	}
}

// Login 管理员登录
func (s *AdminService) Login(username, password string) (*models.Admin, error) {
	var admin models.Admin
	result := s.DB.Where("username = ? AND password = ?", username, password).First(&admin)
	if result.Error != nil {
		return nil, errors.New("用户名或密码错误")
	}
	return &admin, nil
}

// GetAdminByID 根据ID获取管理员
func (s *AdminService) GetAdminByID(id uint) (*models.Admin, error) {
	admin, err := FindOne[models.Admin](s.DB, id)
	return &admin, err
}

// UpdatePassword 更新密码
func (s *AdminService) UpdatePassword(id uint, oldPassword, newPassword string) error {
	var admin models.Admin
	s.DB.First(&admin, id)
	if admin.Password != oldPassword {
		return errors.New("旧密码不正确")
	}
	admin.Password = newPassword
	return Update(s.DB, &admin)
}