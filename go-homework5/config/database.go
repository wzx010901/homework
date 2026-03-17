package config

import (
	"blog-system/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() error {
	var err error
	dbPath := "blog.db"
	if GlobalConfig != nil && GlobalConfig.Database.Path != "" {
		dbPath = GlobalConfig.Database.Path
	}

	models.DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("连接数据库失败：%w", err)
	}

	if err := models.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	); err != nil {
		return fmt.Errorf("数据库迁移失败：%w", err)
	}

	fmt.Println("数据库初始化成功！")
	return nil
}

func GetDB() *gorm.DB {
	return models.DB
}
