package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary() *cloudinary.Cloudinary {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		log.Fatal("CLOUDINARY_URL is not set in environment variables")
	}

	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	return cld
}
