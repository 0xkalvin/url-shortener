package database

import (
	log "github.com/0xkalvin/url-shortener/logger"

	"os"

	redis "github.com/go-redis/redis/v8"
)

// InitializeRedis sets up a connection with redis server
func InitializeRedis() *redis.Client {
	options := &redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	client := redis.NewClient(options)

	logger := log.GetLogger()

	logger.Info("Redis is connected")

	return client
}
