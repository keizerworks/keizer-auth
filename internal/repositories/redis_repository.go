package repositories

import (
	"keizer-auth/internal/database"
	"time"
)

type RedisRepository struct {
	rds *database.RedisService
}

func NewRedisRepository(rds *database.RedisService) *RedisRepository {
	return &RedisRepository{rds: rds}
}

func (rr *RedisRepository) Get(key string) (string, error) {
	value, err := rr.rds.RedisClient.Get(rr.rds.Ctx, key).Result()
	return value, err
}

// set a key's value with expiration
func (rr *RedisRepository) Set(key string, value string, expiration time.Duration) error {
	err := rr.rds.RedisClient.Set(rr.rds.Ctx, key, value, expiration).Err()
	return err
}

func (rr *RedisRepository) TTL(key string) (time.Duration, error) {
	result, err := rr.rds.RedisClient.TTL(rr.rds.Ctx, key).Result()
	return result, err
}
