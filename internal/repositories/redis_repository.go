package repositories

import (
	"time"

	"keizer-auth/internal/database"
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
func (rr *RedisRepository) Set(key string, value interface{}, expiration time.Duration) error {
	err := rr.rds.RedisClient.Set(rr.rds.Ctx, key, value, expiration).Err()
	return err
}

func (rr *RedisRepository) Expire(key string, expiration time.Duration) error {
	return rr.rds.RedisClient.Expire(rr.rds.Ctx, key, expiration).Err()
}

func (rr *RedisRepository) TTL(key string) (time.Duration, error) {
	result, err := rr.rds.RedisClient.TTL(rr.rds.Ctx, key).Result()
	return result, err
}
