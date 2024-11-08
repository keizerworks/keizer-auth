package app

import (
	"sync"

	"keizer-auth-api/internal/database"
	"keizer-auth-api/internal/repositories"
	"keizer-auth-api/internal/services"
)

type Container struct {
	DB           database.Service
	RedisService database.RedisService
	AuthService  *services.AuthService
}

var (
	container *Container
	once      sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.New()
		gormDB := database.GetDB()
		rds := database.NewRedisClient()

		userRepo := repositories.NewUserRepository(gormDB)
		authService := services.NewAuthService(userRepo, rds)

		container = &Container{
			DB:           db,
			AuthService:  authService,
			RedisService: *rds,
		}
	})
	return container
}

func (c *Container) Cleanup() error {
	return c.DB.Close()
}
