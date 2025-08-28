package models

import (
	"time"
	"gorm.io/gorm"
)

// Order 订单模型

type Order struct {
	gorm.Model
	OrderNo     string    `gorm:"size:50;not null;unique" json:"order_no"`
	CustomerID  uint      `gorm:"index;not null" json:"customer_id"`
	ProductID   uint      `gorm:"index;not null" json:"product_id"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	TotalPrice  float64   `gorm:"not null" json:"total_price"`
	Status      string    `gorm:"size:20;not null;default:'pending'" json:"status"`
	OrderTime   time.Time `gorm:"not null" json:"order_time"`
	Customer    Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	Product     Product   `gorm:"foreignKey:ProductID" json:"product"`
}

// 表名重命名
func (Order) TableName() string {
	return "order"
}