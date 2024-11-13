package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			// TODO: handle cors domain
			return strings.Contains(origin, "localhost")
		},
		AllowCredentials: true,
	}))

	s.Get("/health", s.healthHandler)

	api := s.Group("/api")

	// auth handlers
	auth := api.Group("/auth")
	auth.Post("/sign-up", s.controllers.Auth.SignUp)
	auth.Post("/sign-in", s.controllers.Auth.SignIn)
	auth.Post("/verify-otp", s.controllers.Auth.VerifyOTP)
	auth.Get("/verify-token", s.controllers.Auth.VerifyTokenHandler)

	s.Static("/", "./web/dist")
	s.Static("*", "./web/dist/index.html")
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.container.DB.Health())
}
