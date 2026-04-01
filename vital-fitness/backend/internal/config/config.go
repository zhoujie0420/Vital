package config

import (
	"fmt"

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

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	return &config
}
