package services

import (
	"github.com/supermarketmanager/database"
	"github.com/supermarketmanager/models"
	"gorm.io/gorm"
)

// ProductService 商品服务

type ProductService struct {
	DB *gorm.DB
}

// NewProductService 创建商品服务实例
func NewProductService() *ProductService {
	return &ProductService{
		DB: database.GetDB(),
	}
}

// GetAllProducts 获取所有商品，包括关联的分类和供应商信息
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := s.DB.Preload("Category").Preload("Supplier").Find(&products)
	return products, result.Error
}

// GetProductByID 根据ID获取商品，包括关联的分类和供应商信息
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := s.DB.Preload("Category").Preload("Supplier").First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

// CreateProduct 创建商品
func (s *ProductService) CreateProduct(product *models.Product) error {
	return Create(s.DB, product)
}

// UpdateProduct 更新商品
func (s *ProductService) UpdateProduct(product *models.Product) error {
	return Update(s.DB, product)
}

// DeleteProduct 删除商品
func (s *ProductService) DeleteProduct(id uint) error {
	return Delete[models.Product](s.DB, id)
}

// QueryProducts 根据条件查询商品
func (s *ProductService) QueryProducts(query map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	db := s.DB.Preload("Category").Preload("Supplier")
	
	// 根据查询条件构建查询
	if name, ok := query["name"].(string); ok && name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if categoryID, ok := query["category_id"].(uint); ok && categoryID > 0 {
		db = db.Where("category_id = ?", categoryID)
	}
	if supplierID, ok := query["supplier_id"].(uint); ok && supplierID > 0 {
		db = db.Where("supplier_id = ?", supplierID)
	}
	if minPrice, ok := query["min_price"].(float64); ok && minPrice > 0 {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice, ok := query["max_price"].(float64); ok && maxPrice > 0 {
		db = db.Where("price <= ?", maxPrice)
	}
	
	result := db.Find(&products)
	return products, result.Error
}