package database

import (
	redis "github.com/go-redis/redis/v8"
)

// InitializeRedis sets up a connection with redis server
func InitializeRedis() *redis.Client {
	options := &redis.Options{
		Addr:     "redis:6379",
		Password: "",
	}

	client := redis.NewClient(options)

	return client
}
