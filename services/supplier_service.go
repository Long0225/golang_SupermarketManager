package services

import (
	"github.com/supermarketmanager/database"
	"github.com/supermarketmanager/models"
	"gorm.io/gorm"
)

// SupplierService 供应商服务

type SupplierService struct {
	DB *gorm.DB
}

// NewSupplierService 创建供应商服务实例
func NewSupplierService() *SupplierService {
	return &SupplierService{
		DB: database.GetDB(),
	}
}

// GetAllSuppliers 获取所有供应商
func (s *SupplierService) GetAllSuppliers() ([]models.Supplier, error) {
	return FindAll[models.Supplier](s.DB)
}

// GetSupplierByID 根据ID获取供应商
func (s *SupplierService) GetSupplierByID(id uint) (*models.Supplier, error) {
	supplier, err := FindOne[models.Supplier](s.DB, id)
	return &supplier, err
}

// CreateSupplier 创建供应商
func (s *SupplierService) CreateSupplier(supplier *models.Supplier) error {
	return Create(s.DB, supplier)
}

// UpdateSupplier 更新供应商
func (s *SupplierService) UpdateSupplier(supplier *models.Supplier) error {
	return Update(s.DB, supplier)
}

// DeleteSupplier 删除供应商
func (s *SupplierService) DeleteSupplier(id uint) error {
	return Delete[models.Supplier](s.DB, id)
}