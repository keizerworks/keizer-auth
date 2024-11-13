package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nrednav/cuid2"
)

const SessionExpiresIn = 30 * 24 * time.Hour

func GenerateSessionID() (string, error) {
	generate, err := cuid2.Init(
		cuid2.WithLength(15),
	)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return generate(), nil
}

func SetSessionCookie(c *fiber.Ctx, sessionID string) {
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(SessionExpiresIn),
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
		// TODO: handle domain
		Domain: "localhost",
		Path:   "/",
	})
}

func GetSessionCookie(c *fiber.Ctx) string {
	return c.Cookies("session_id", "")
}
