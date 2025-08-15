package main

import (
	"log"

	"backend/config"
)

func main() {
	cfg := config.LoadConfig()

	db := config.InitDB(cfg)
	redisClient := config.InitRedis(cfg)
	cloud := config.InitCloudinary(cfg)

	log.Println(db, redisClient, cloud)

	// http
}
