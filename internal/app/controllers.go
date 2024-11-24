package app

import "keizer-auth/internal/controllers"

type ServerControllers struct {
	Auth        *controllers.AuthController
	Account     *controllers.AccountController
	Application *controllers.ApplicationController
}

func GetControllers(container *Container) *ServerControllers {
	return &ServerControllers{
		Auth:        controllers.NewAuthController(container.AuthService, container.SessionService),
		Account:     controllers.NewAccountController(container.AccountService),
		Application: controllers.NewApplicationController(container.ApplicationService),
	}
}
