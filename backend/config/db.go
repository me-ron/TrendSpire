package config

import (
	"backend/internal/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE") 

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, dbName, port, sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	return db
}

func RunMigrations() {
	err := DB.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Comment{},
		&entity.Like{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}