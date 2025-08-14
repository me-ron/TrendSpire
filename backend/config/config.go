package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Redis
	RedisHost string
	RedisPort string
	RedisPass string
	RedisDB   int

	// Cloudinary
	CloudinaryURL string

	// JWT
	JWTSecret string
	JWTExpiry time.Duration
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		AppPort: viper.GetString("APP_PORT"),

		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBSSLMode:  viper.GetString("DB_SSLMODE"),

		RedisHost: viper.GetString("REDIS_HOST"),
		RedisPort: viper.GetString("REDIS_PORT"),
		RedisPass: viper.GetString("REDIS_PASSWORD"),
		RedisDB:   viper.GetInt("REDIS_DB"),

		CloudinaryURL: viper.GetString("CLOUDINARY_URL"),

		JWTSecret: viper.GetString("JWT_SECRET"),
		JWTExpiry: viper.GetDuration("JWT_EXPIRY") * time.Minute,
	}

	return cfg
}
