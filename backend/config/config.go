package config

import (
	"backend/pkg/utils"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Env  string
		Port string
	}

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}

	Cloudinary struct {
		CloudName string
		APIKey    string
		APISecret string
	}

	JWT struct {
		AccessSecret  string
		RefreshSecret string
		Issuer        string
		AccessExpire  time.Duration
		RefreshExpire time.Duration
	}

	Log struct {
		Level string
	}
}

func LoadConfig() *Config {
	// Load .env if present
	_ = godotenv.Load()

	viper.AutomaticEnv()

	cfg := &Config{}

	// App
	cfg.App.Name = viper.GetString("APP_NAME")
	cfg.App.Env = viper.GetString("APP_ENV")
	cfg.App.Port = viper.GetString("APP_PORT")

	// Database
	cfg.DB.Host = viper.GetString("DB_HOST")
	cfg.DB.Port = viper.GetString("DB_PORT")
	cfg.DB.User = viper.GetString("DB_USER")
	cfg.DB.Password = viper.GetString("DB_PASSWORD")
	cfg.DB.Name = viper.GetString("DB_NAME")
	cfg.DB.SSLMode = viper.GetString("DB_SSLMODE")

	// Redis
	cfg.Redis.Host = viper.GetString("REDIS_HOST")
	cfg.Redis.Port = viper.GetString("REDIS_PORT")
	cfg.Redis.Password = viper.GetString("REDIS_PASSWORD")
	cfg.Redis.DB = viper.GetInt("REDIS_DB")

	// Cloudinary
	cfg.Cloudinary.CloudName = viper.GetString("CLOUDINARY_CLOUD_NAME")
	cfg.Cloudinary.APIKey = viper.GetString("CLOUDINARY_API_KEY")
	cfg.Cloudinary.APISecret = viper.GetString("CLOUDINARY_API_SECRET")

	// JWT
	cfg.JWT.AccessSecret = viper.GetString("JWT_ACCESS_SECRET")
	cfg.JWT.RefreshSecret = viper.GetString("JWT_REFRESH_SECRET")
	cfg.JWT.Issuer = viper.GetString("JWT_ISSUER")

	accessExpire, err := utils.ParseDuration(viper.GetString("JWT_ACCESS_EXPIRE"))
	if err != nil {
		log.Fatal("invalid JWT_ACCESS_EXPIRE format")
	}
	cfg.JWT.AccessExpire = accessExpire

	refreshExpire, err := utils.ParseDuration(viper.GetString("JWT_REFRESH_EXPIRE"))
	if err != nil {
		log.Fatal("invalid JWT_REFRESH_EXPIRE format")
	}
	cfg.JWT.RefreshExpire = refreshExpire

	// Logging
	cfg.Log.Level = viper.GetString("LOG_LEVEL")

	return cfg
}
