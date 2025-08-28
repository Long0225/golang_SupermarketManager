package models

import "gorm.io/gorm"

// Admin 管理员模型，对应Java中的Admin实体

type Admin struct {
	gorm.Model
	Username string `gorm:"size:50;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null" json:"password"`
	Nickname string `gorm:"size:50" json:"nickname"`
	Note     string `gorm:"size:255" json:"note"`
}

// 表名重命名
func (Admin) TableName() string {
	return "admin"
}
