package services

import (
	"github.com/supermarketmanager/database"
	"github.com/supermarketmanager/models"
	"gorm.io/gorm"
)

// OrderService 订单服务

type OrderService struct {
	DB *gorm.DB
}

// NewOrderService 创建订单服务实例
func NewOrderService() *OrderService {
	return &OrderService{
		DB: database.GetDB(),
	}
}

// GetAllOrders 获取所有订单，包括关联的客户和商品信息
func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := s.DB.Preload("Customer").Preload("Product").Find(&orders)
	return orders, result.Error
}

// GetOrderByID 根据ID获取订单，包括关联的客户和商品信息
func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	result := s.DB.Preload("Customer").Preload("Product").First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(order *models.Order) error {
	return Create(s.DB, order)
}

// UpdateOrder 更新订单
func (s *OrderService) UpdateOrder(order *models.Order) error {
	return Update(s.DB, order)
}

// DeleteOrder 删除订单
func (s *OrderService) DeleteOrder(id uint) error {
	return Delete[models.Order](s.DB, id)
}