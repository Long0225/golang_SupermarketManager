package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// 当前选择区域为空，且错误提示表明 “log” 包导入未使用，由于当前代码中未发现 “log” 包导入，无需修改此选择区域内容。
)

var DB *gorm.DB

func InitDB() error {
	user := "root"
	password := "0225@long" // 直接使用原始密码
	host := "localhost"
	port := "3306"
	database := "supermarketmanager"
	// 连接数据库的配置信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移表结构
	// 注意：在实际生产环境中，建议手动管理数据库迁移

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
