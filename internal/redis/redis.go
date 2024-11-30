package redis

import (
	"context"
	"os"

	"github.com/codebyaadi/rss-feed-agg/pkg/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var Client *redis.Client

func InitRedis() error {
	url := os.Getenv("REDIS_URL")
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
		logger.Logger.Info("defaulting to addr %s", zap.String("addr", addr))
	}
	password := os.Getenv("REDIS_PASSWORD")
	db := 0

	if url != "" {
		opt, err := redis.ParseURL(url)
		if err != nil {
			logger.Logger.Fatal("failed to parse Redis URL: %v", zap.Error(err))
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
		logger.Logger.Fatal("failed to connect to redis: %v", zap.Error(err))
		Client.Close()
		return err
	}

	logger.Logger.Info("successfully connected to Redis")
	return nil
}

func CloseRedis() error {
	return Client.Close()
}
