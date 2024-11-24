package app

import (
	"keizer-auth/internal/middlewares"
)

type ServerMiddlewares struct {
	Auth *middlewares.AuthMiddleware
}

func GetMiddlewares(container *Container) *ServerMiddlewares {
	return &ServerMiddlewares{
		Auth: middlewares.NewAuthMiddleware(container.SessionService),
	}
}
