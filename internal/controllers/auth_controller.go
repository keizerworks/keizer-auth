package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController struct{}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	// Logic to handle login
	// For example, parse the request and validate user credentials
	return c.SendString("Login successful!")
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	// Logic to handle user registration
	return c.SendString("User registered!")
}
