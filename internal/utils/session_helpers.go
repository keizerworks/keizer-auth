package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nrednav/cuid2"
)

const SessionExpiresIn = 30 * 24 * time.Hour

func GenerateSessionID() string {
	return cuid2.Generate()
}

func SetSessionCookie(c *fiber.Ctx, sessionID string) {
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(SessionExpiresIn),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteNoneMode,
		// TODO: handle domain
	})
}

func GetSessionCookie(c *fiber.Ctx) string {
	return c.Cookies("session_id", "")
}
