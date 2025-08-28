package models

import (
	"time"
	"gorm.io/gorm"
)

// SystemLog 系统日志模型

type SystemLog struct {
	gorm.Model
	Operator    string    `gorm:"size:50;not null" json:"operator"`
	OperateTime time.Time `gorm:"not null" json:"operate_time"`
	Action      string    `gorm:"size:50;not null" json:"action"`
	Details     string    `gorm:"size:255" json:"details"`
	IP          string    `gorm:"size:20" json:"ip"`
}

// 表名重命名
func (SystemLog) TableName() string {
	return "system_log"
}