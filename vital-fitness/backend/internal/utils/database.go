package utils

import (
	"fmt"
	"time"

	"vital-fitness/backend/internal/config"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDatabase 初始化数据库连接（带重试）
func InitDatabase(cfg config.Database) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	var err error
	// 最多重试 30 次，每次间隔 2 秒，等 MySQL 启动
	for i := 0; i < 30; i++ {
		db, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
		if err == nil {
			break
		}
		GetLogger().Warn("waiting for database...", zap.Int("attempt", i+1), zap.Error(err))
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		GetLogger().Fatal("failed to connect database after retries", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		GetLogger().Fatal("failed to get sql db", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	GetLogger().Info("database connected successfully")
	return db
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return db
}
