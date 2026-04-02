package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Config 应用配置结构体
type Config struct {
	ServerPort string   `mapstructure:"SERVER_PORT"`
	Mode       string   `mapstructure:"MODE"`
	LogLevel   string   `mapstructure:"LOG_LEVEL"`
	Database   Database `mapstructure:"DATABASE"`
	JWT        JWT      `mapstructure:"JWT"`
}

// Database 数据库配置
type Database struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

// JWT JWT配置
type JWT struct {
	Secret string `mapstructure:"JWT_SECRET"`
	Expire int    `mapstructure:"JWT_EXPIRE"`
}

// LoadConfig 加载配置
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("../../configs")

	// 设置默认值
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("MODE", "debug")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("DATABASE.DB_HOST", "localhost")
	viper.SetDefault("DATABASE.DB_PORT", "3306")
	viper.SetDefault("DATABASE.DB_USER", "root")
	viper.SetDefault("DATABASE.DB_NAME", "vital_fitness")
	viper.SetDefault("JWT.JWT_EXPIRE", 24)

	// 读取配置文件（不存在也不报错，用默认值和环境变量）
	_ = viper.ReadInConfig()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	// 环境变量覆盖（Docker 部署时使用）
	envOverride(&config.ServerPort, "SERVER_PORT")
	envOverride(&config.Mode, "MODE")
	envOverride(&config.LogLevel, "LOG_LEVEL")
	envOverride(&config.Database.Host, "DB_HOST")
	envOverride(&config.Database.Port, "DB_PORT")
	envOverride(&config.Database.User, "DB_USER")
	envOverride(&config.Database.Password, "DB_PASSWORD")
	envOverride(&config.Database.Name, "DB_NAME")
	envOverride(&config.JWT.Secret, "JWT_SECRET")
	if v := os.Getenv("JWT_EXPIRE"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			config.JWT.Expire = n
		}
	}

	return &config
}

func envOverride(target *string, key string) {
	if v := os.Getenv(key); v != "" {
		*target = v
	}
}
