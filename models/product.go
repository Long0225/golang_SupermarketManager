package models

import "gorm.io/gorm"

// Product 商品模型，对应Java中的Product实体

type Product struct {
	gorm.Model
	Code       string  `gorm:"size:50;not null;unique" json:"code"`
	Name       string  `gorm:"size:100;not null" json:"name"`
	Price      float64 `gorm:"not null" json:"price"`
	Stock      int     `gorm:"not null" json:"stock"`
	CategoryID uint    `gorm:"index;not null" json:"category_id"`
	SupplierID uint    `gorm:"index;not null" json:"supplier_id"`
	Desc       string  `gorm:"size:255" json:"desc"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID" json:"supplier"`
}

// 表名重命名
func (Product) TableName() string {
	return "product"
}