package app

import (
	"sync"

	"keizer-auth-api/internal/database"
	"keizer-auth-api/internal/repositories"
	"keizer-auth-api/internal/services"
)

type Container struct {
	DB          database.Service
	AuthService *services.AuthService
}

var (
	container *Container
	once      sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.New()
		gormDB := database.GetDB()

		userRepo := repositories.NewUserRepository(gormDB)
		authService := services.NewAuthService(userRepo)

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
