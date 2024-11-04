package middlewares

import (
	"keizer-auth-api/internal/database"
	"keizer-auth-api/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

var allowedDomains = []string{
	"https://dashboard.auth.keizerworks.com",
	"http://localhost:3000",
}

// Middleware to validate request origin
func OriginValidationMiddleware(c *fiber.Ctx) error {
	// Allow if running on localhost
	if c.Hostname() == "localhost" || c.Hostname() == "127.0.0.1" {
		return c.Next()
	}

	origin := c.Get("Origin")

	for _, domain := range allowedDomains {
		if domain == origin {
			return c.Next()
		}
	}

	db := database.GetDB()
	domainRepository := repositories.NewDomainRepository(db)

	if _, err := domainRepository.GetActiveDomain(origin); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Origin not allowed",
		})
	}

	return c.Next()
}
