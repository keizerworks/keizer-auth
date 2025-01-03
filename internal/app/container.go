package app

import (
	"keizer-auth/internal/database"
	"keizer-auth/internal/repositories"
	"keizer-auth/internal/services"
	"sync"
)

type Container struct {
	DB                 database.Service
	AuthService        *services.AuthService
	SessionService     *services.SessionService
	EmailService       *services.EmailService
	AccountService     *services.AccountService
	ApplicationService *services.ApplicationService
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
		accountRepo := repositories.NewAccountRepository(gormDB)
		applicationRepo := repositories.NewApplicationRepository(gormDB)
		userAccountRepo := repositories.NewUserAccountRepository(gormDB)
		redisRepo := repositories.NewRedisRepository(rds)

		authService := services.NewAuthService(userRepo, redisRepo)
		sessionService := services.NewSessionService(redisRepo, userRepo)
		accountService := services.NewAccountService(accountRepo, userAccountRepo)
		applicationService := services.NewApplicationService(applicationRepo, accountRepo)

		container = &Container{
			DB:                 db,
			AuthService:        authService,
			SessionService:     sessionService,
			AccountService:     accountService,
			ApplicationService: applicationService,
		}
	})
	return container
}

func (c *Container) Cleanup() error {
	return c.DB.Close()
}
