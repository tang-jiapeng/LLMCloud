package database

import (
	"fmt"
	"llmcloud/config"
	"llmcloud/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfigInstance.Database.User,
		config.AppConfigInstance.Database.Password,
		config.AppConfigInstance.Database.Host,
		config.AppConfigInstance.Database.Port,
		config.AppConfigInstance.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 自动迁移
	if err := db.AutoMigrate(&model.User{}, &model.File{}); err != nil {
		return nil, err
	}

	return db, nil
}
