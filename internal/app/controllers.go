package app

import "keizer-auth/internal/controllers"

type ServerControllers struct {
	Auth *controllers.AuthController
}

func GetControllers(container *Container) *ServerControllers {
	return &ServerControllers{
		Auth: controllers.NewAuthController(container.AuthService, container.SessionService),
	}
}
