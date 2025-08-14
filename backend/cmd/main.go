package main

import (
	"log"

	"backend/config"
)

func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	// Init services
	db := config.InitDB(cfg)
	redisClient := config.InitRedis(cfg)
	cloud := config.InitCloudinary(cfg)

	log.Println(db, redisClient, cloud)

	// http
}
