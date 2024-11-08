package repositories

import "keizer-auth-api/internal/database"

type RedisRepository struct {
	rdb *database.RedisService
}

func NewRedisRepository(rds *database.RedisService) *RedisRepository {
	return &RedisRepository{rdb: rds}
}
