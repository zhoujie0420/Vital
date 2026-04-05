package main

import (
	"fmt"
	"os"

	"vital-fitness/backend/internal/config"
	"vital-fitness/backend/internal/middleware"
	"vital-fitness/backend/internal/router"
	"vital-fitness/backend/internal/utils"
)

// @title Vital Fitness API
// @version 1.0
// @description Vital Fitness身体状态记录系统API
// @termsOfService https://example.com/terms/

// @contact.name API Support
// @contact.url https://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	// 初始化配置
	cfg := config.LoadConfig()

	// 初始化日志
	utils.InitLogger(cfg.LogLevel)

	// 初始化数据库
	db := utils.InitDatabase(cfg.Database)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 自动迁移（添加新字段）
	db.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS open_id VARCHAR(100) UNIQUE AFTER deleted_at")

	// 初始化JWT
	middleware.InitJWT(cfg.JWT)

	// 设置路由
	r := router.SetupRouter(cfg.Mode, cfg.WeChat)

	// 启动服务
	port := fmt.Sprintf(":%s", cfg.ServerPort)
	if port == ":" {
		port = ":8080"
	}

	fmt.Printf("Starting server on port %s\n", port)
	if err := r.Run(port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
