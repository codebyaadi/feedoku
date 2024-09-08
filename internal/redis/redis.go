package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis() error {
	url := os.Getenv("REDIS_URL")
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
		log.Printf("defaulting to addr %s", addr)
	}
	password := os.Getenv("REDIS_PASSWORD")
	db := 0

	if url != "" {
		opt, err := redis.ParseURL(url)
		if err != nil {
			log.Fatalf("failed to parse Redis URL: %v", err)
		}
		Client = redis.NewClient(opt)
	} else {
		Client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
	}

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
		Client.Close()
		return err
	}

	log.Println("successfully connected to Redis")
	return nil
}

func CloseRedis() error {
	return Client.Close()
}
