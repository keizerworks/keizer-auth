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
	auth.Get("/profile", s.middlewars.Auth.Authorize, s.controllers.Auth.Profile)

	// accounts handlers
	accounts := api.Group("/accounts", s.middlewars.Auth.Authorize)
	accounts.Post("/", s.controllers.Account.Create)
	accounts.Get("/", s.controllers.Account.Get)

	// applications handlers
	applications := accounts.Group("/:accountId<guid>/applications")
	applications.Post("/", s.controllers.Application.Create)
	applications.Get("/", s.controllers.Application.Get)

	s.Static("/", "./web/dist")
	s.Static("*", "./web/dist/index.html")
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.container.DB.Health())
}
