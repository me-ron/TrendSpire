package main

import (
	"log"

	"backend/config"
	"backend/internal/entity"
)

func main() {
	cfg := config.LoadConfig()

	db := config.InitDB(cfg)

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatalf("❌ Failed to enable uuid-ossp extension: %v", err)
	}

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Like{},
		&entity.Comment{},
	); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	if db.Migrator().HasColumn(&entity.Post{}, "username") {
		if err := db.Migrator().DropColumn(&entity.Post{}, "username"); err != nil {
			log.Fatalf("❌ Failed to drop username column: %v", err)
		}
	}
	if db.Migrator().HasColumn(&entity.Like{}, "username") {
		if err := db.Migrator().DropColumn(&entity.Like{}, "username"); err != nil {
			log.Fatalf("❌ Failed to drop username column: %v", err)
		}
	}
	if db.Migrator().HasColumn(&entity.Comment{}, "username") {
		if err := db.Migrator().DropColumn(&entity.Comment{}, "username"); err != nil {
			log.Fatalf("❌ Failed to drop username column: %v", err)
		}
	}

	log.Println("✅ Migrations completed successfully")
}
