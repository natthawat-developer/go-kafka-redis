package redis

import (
	"context"
	"go-kafka-redis/pkg/logger" 
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(addr string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func SaveToRedis(key string, value string) {
	ctx := context.Background()
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to save to Redis: %s", err)
	} else {
		logger.InfoLogger.Printf("Saved to Redis: %s -> %s", key, value)
	}
}

func CloseRedis() {
	redisClient.Close()
}
