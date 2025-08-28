package services

import (
	"gorm.io/gorm"
)

// BaseService 基础服务接口

type BaseService interface {
	FindAll() ([]interface{}, error)
	FindOne(id uint) (interface{}, error)
	Create(entity interface{}) error
	Update(entity interface{}) error
	Delete(id uint) error
}

// 基础服务的通用实现

func FindAll[T any](db *gorm.DB) ([]T, error) {
	var entities []T
	result := db.Find(&entities)
	return entities, result.Error
}

func FindOne[T any](db *gorm.DB, id uint) (T, error) {
	var entity T
	result := db.First(&entity, id)
	return entity, result.Error
}

func Create[T any](db *gorm.DB, entity T) error {
	result := db.Create(entity)
	return result.Error
}

func Update[T any](db *gorm.DB, entity T) error {
	result := db.Save(entity)
	return result.Error
}

func Delete[T any](db *gorm.DB, id uint) error {
	var entity T
	result := db.Delete(&entity, id)
	return result.Error
}