package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *RedisService

type RedisService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewRedisClient() *RedisService {
	if rdb != nil {
		return rdb
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:   0,
	})
	ctx := context.Background()

	rdb = &RedisService{redisClient: redisClient, ctx: ctx}

	return rdb
}
