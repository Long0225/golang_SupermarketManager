package services

import (
	"github.com/supermarketmanager/database"
	"github.com/supermarketmanager/models"
	"gorm.io/gorm"
)

// CustomerService 客户服务

type CustomerService struct {
	DB *gorm.DB
}

// NewCustomerService 创建客户服务实例
func NewCustomerService() *CustomerService {
	return &CustomerService{
		DB: database.GetDB(),
	}
}

// GetAllCustomers 获取所有客户
func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	return FindAll[models.Customer](s.DB)
}

// GetCustomerByID 根据ID获取客户
func (s *CustomerService) GetCustomerByID(id uint) (*models.Customer, error) {
	customer, err := FindOne[models.Customer](s.DB, id)
	return &customer, err
}

// CreateCustomer 创建客户
func (s *CustomerService) CreateCustomer(customer *models.Customer) error {
	return Create(s.DB, customer)
}

// UpdateCustomer 更新客户
func (s *CustomerService) UpdateCustomer(customer *models.Customer) error {
	return Update(s.DB, customer)
}

// DeleteCustomer 删除客户
func (s *CustomerService) DeleteCustomer(id uint) error {
	return Delete[models.Customer](s.DB, id)
}