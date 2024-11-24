package middlewares

import (
	"keizer-auth/internal/constants"
	"keizer-auth/internal/models"
	"keizer-auth/internal/services"
	"keizer-auth/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	sessionService *services.SessionService
}

func NewAuthMiddleware(
	ss *services.SessionService,
) *AuthMiddleware {
	return &AuthMiddleware{sessionService: ss}
}

func (self *AuthMiddleware) Authorize(c *fiber.Ctx) error {
	sessionID := utils.GetSessionCookie(c)
	if sessionID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var user models.User
	err := self.sessionService.GetSession(sessionID, &user)
	if err != nil {
		log.Printf("Session validation error: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	c.Locals(constants.UserContextKey, &user)
	return c.Next()
}
