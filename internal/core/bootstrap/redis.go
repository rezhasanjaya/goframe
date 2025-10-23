package bootstrap

import (
	"context"
	"log"

	"goframe/internal/core/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func InitRedis(cfg *config.Config) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost,
	})
	if _, err := RedisClient.Ping(RedisCtx).Result(); err != nil {
		log.Printf("⚠️ Redis not connected: %v", err)
	} else {
		log.Println("✅ Redis connected")
	}
}
