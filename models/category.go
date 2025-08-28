package models

import "gorm.io/gorm"

// Category 分类模型，对应Java中的Category实体

type Category struct {
	gorm.Model
	Name string `gorm:"size:50;not null;unique" json:"name"`
	Desc string `gorm:"size:255" json:"desc"`
}

// 表名重命名
func (Category) TableName() string {
	return "category"
}