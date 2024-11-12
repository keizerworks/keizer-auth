package app

import (
	"sync"

	"keizer-auth/internal/database"
	"keizer-auth/internal/repositories"
	"keizer-auth/internal/services"
)

type Container struct {
	DB             database.Service
	AuthService    *services.AuthService
	SessionService *services.SessionService
	EmailService   *services.EmailService
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
		sessionService := services.NewSessionService(redisRepo, userRepo)

		container = &Container{
			DB:             db,
			AuthService:    authService,
			SessionService: sessionService,
		}
	})
	return container
}

func (c *Container) Cleanup() error {
	return c.DB.Close()
}
