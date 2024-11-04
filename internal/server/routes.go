package server

import (
	"keizer-auth-api/internal/handlers"
	"keizer-auth-api/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.Get("/health", s.healthHandler)

	api := s.Group("/api", middlewares.OriginValidationMiddleware)

	// auth handlers
	auth := api.Group("/auth")
	auth.Post("/register", s.controllers.Auth.Register)
	auth.Post("/login", s.controllers.Auth.Login)

	handlers.RegisterAuthHandlers(api)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.container.DB.Health())
}
