package main

import (
	"backend/config"
)

func main() {
	config.LoadEnv()
	db := config.InitPostgres()
	redis := config.InitRedis()
	cld := config.InitCloudinary()

	// TODO: Migrations here using db.AutoMigrate(...)
	_ = db
	_ = redis
	_ = cld
}
