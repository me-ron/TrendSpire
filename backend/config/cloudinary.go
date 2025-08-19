package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary(cfg *Config) *cloudinary.Cloudinary {
	var cld *cloudinary.Cloudinary
	var err error

	if cfg.Cloudinary.CloudName != "" && cfg.Cloudinary.APIKey != "" && cfg.Cloudinary.APISecret != "" {
		// Fallback to individual credentials
		cld, err = cloudinary.NewFromParams(cfg.Cloudinary.CloudName, cfg.Cloudinary.APIKey , cfg.Cloudinary.APISecret)
	} else {
		log.Fatal("❌ Missing Cloudinary configuration (provide CLOUDINARY_URL or all of CloudName, CloudAPIKey, CloudAPISecret)")
	}

	if err != nil {
		log.Fatalf("❌ Failed to initialize Cloudinary: %v", err)
	}

	log.Println("✅ Connected to Cloudinary")
	return cld
}
