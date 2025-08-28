package services

import (
	"github.com/supermarketmanager/database"
	"github.com/supermarketmanager/models"
	"gorm.io/gorm"
)

// CategoryService 分类服务

type CategoryService struct {
	DB *gorm.DB
}

// NewCategoryService 创建分类服务实例
func NewCategoryService() *CategoryService {
	return &CategoryService{
		DB: database.GetDB(),
	}
}

// GetAllCategories 获取所有分类
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return FindAll[models.Category](s.DB)
}

// GetCategoryByID 根据ID获取分类
func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	category, err := FindOne[models.Category](s.DB, id)
	return &category, err
}

// CreateCategory 创建分类
func (s *CategoryService) CreateCategory(category *models.Category) error {
	return Create(s.DB, category)
}

// UpdateCategory 更新分类
func (s *CategoryService) UpdateCategory(category *models.Category) error {
	return Update(s.DB, category)
}

// DeleteCategory 删除分类
func (s *CategoryService) DeleteCategory(id uint) error {
	return Delete[models.Category](s.DB, id)
}