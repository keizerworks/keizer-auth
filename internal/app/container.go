package app

import (
	"sync"

	"keizer-auth-api/internal/database"
	"keizer-auth-api/internal/repositories"
	"keizer-auth-api/internal/services"
)

type Container struct {
	DB             database.Service
	AuthService    *services.AuthService
	SessionService *services.SessionService
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
		redisRepo := repositories.NewRedisRepository(rds)
		authService := services.NewAuthService(userRepo, redisRepo)

		container = &Container{
			DB:          db,
			AuthService: authService,
		}
	})
	return container
}

func (c *Container) Cleanup() error {
	return c.DB.Close()
}
