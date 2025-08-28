package models

import "gorm.io/gorm"

// Supplier 供应商模型，对应Java中的Supplier实体

type Supplier struct {
	gorm.Model
	Code       string `gorm:"size:50;not null;unique" json:"code"`
	Name       string `gorm:"size:100;not null" json:"name"`
	Contact    string `gorm:"size:50" json:"contact"`
	Phone      string `gorm:"size:20" json:"phone"`
	Address    string `gorm:"size:255" json:"address"`
}

// 表名重命名
func (Supplier) TableName() string {
	return "supplier"
}