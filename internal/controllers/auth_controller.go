package controllers

import (
	"keizer-auth-api/internal/services"
	"keizer-auth-api/internal/validators"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{authService: as}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	return c.SendString("Login successful!")
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	user := new(validators.UserRegister)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if valid, errors := user.Validate(); !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	if err := ac.authService.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.SendString("User registered!")
}
