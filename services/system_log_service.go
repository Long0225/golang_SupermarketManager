package services

import (
	"time"
	"github.com/supermarketmanager/database"
	"github.com/supermarketmanager/models"
	"gorm.io/gorm"
)

// SystemLogService 系统日志服务

type SystemLogService struct {
	DB *gorm.DB
}

// NewSystemLogService 创建系统日志服务实例
func NewSystemLogService() *SystemLogService {
	return &SystemLogService{
		DB: database.GetDB(),
	}
}

// GetAllSystemLogs 获取所有系统日志
func (s *SystemLogService) GetAllSystemLogs() ([]models.SystemLog, error) {
	var logs []models.SystemLog
	result := s.DB.Order("created_at DESC").Find(&logs)
	return logs, result.Error
}

// GetSystemLogByID 根据ID获取系统日志
func (s *SystemLogService) GetSystemLogByID(id uint) (*models.SystemLog, error) {
	log, err := FindOne[models.SystemLog](s.DB, id)
	return &log, err
}

// CreateSystemLog 创建系统日志
func (s *SystemLogService) CreateSystemLog(log *models.SystemLog) error {
	log.OperateTime = time.Now()
	return Create(s.DB, log)
}