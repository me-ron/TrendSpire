package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary(cfg *Config) *cloudinary.Cloudinary {
	if cfg.CloudinaryURL == "" {
		log.Fatal("❌ CLOUDINARY_URL is not set in configuration")
	}

	cld, err := cloudinary.NewFromURL(cfg.CloudinaryURL)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Cloudinary: %v", err)
	}

	log.Println("✅ Connected to Cloudinary")
	return cld
}
