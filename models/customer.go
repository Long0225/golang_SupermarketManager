package models

import "gorm.io/gorm"

// Customer 顾客模型

type Customer struct {
	gorm.Model
	Name    string `gorm:"size:50;not null" json:"name"`
	Phone   string `gorm:"size:20;unique" json:"phone"`
	Address string `gorm:"size:255" json:"address"`
	Gender  string `gorm:"size:10" json:"gender"`
}

// 表名重命名
func (Customer) TableName() string {
	return "customer"
}