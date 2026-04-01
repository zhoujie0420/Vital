package utils

import (
	"fmt"
	"time"

	"vital-fitness/backend/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

var db *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase(cfg config.Database) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		GetLogger().Fatal("failed to connect database", zap.Error(err))
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		GetLogger().Fatal("failed to get sql db", zap.Error(err))
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	GetLogger().Info("database connected successfully")
	return db
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return db
}