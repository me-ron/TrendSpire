package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *Config) *redis.Client {
	addresses := []string{
		fmt.Sprintf("[%s]:%s", cfg.Redis.Host, cfg.Redis.Port), 
		fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port), 
	}

	var rdb *redis.Client
	var err error

	for _, addr := range addresses {
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = rdb.Ping(ctx).Err()
		if err == nil {
			log.Printf("✅ Connected to Redis at %s\n", addr)
			return rdb
		}

		log.Printf("⚠️ Failed to connect to Redis at %s: %v", addr, err)
	}

	log.Fatalf("❌ Could not connect to Redis on any address: %v", err)
	return nil
}
