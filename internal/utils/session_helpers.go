package utils

import (
	"crypto/rand"
	"encoding/base32"
	"time"

	"github.com/gofiber/fiber/v2"
)

const SessionExpiresIn = 30 * 24 * time.Hour

func GenerateSessionID() (string, error) {
	bytes := make([]byte, 15)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(bytes), nil
}

// func ValidateSession(
// 	sessionID string,
// 	ttl time.Duration,
// ) error {
// 	if ttl < SessionExpiresIn {
// 		if err := sessionService.UpdateSession(sessionID); err != nil {
// 			return nil, err
// 		}
// 	}
// 	return nil
// }

func SetSessionCookie(c *fiber.Ctx, sessionID string) {
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(SessionExpiresIn),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})
}
