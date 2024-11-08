package controllers

import (
	"errors"

	"keizer-auth-api/internal/services"
	"keizer-auth-api/internal/validators"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{authService: as}
}

func (ac *AuthController) SignIn(c *fiber.Ctx) error {
	return c.SendString("Login successful!")
}

func (ac *AuthController) SignUp(c *fiber.Ctx) error {
	user := new(validators.SignUpUser)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if valid, errors := user.Validate(); !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	if err := ac.authService.RegisterUser(user); err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"error": "Input validation error, please check your details",
				})
		}

		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to sign up user"})
	}

	return c.JSON(fiber.Map{"message": "User Signed Up!"})
}

func (ac *AuthController) VerifyOTP(c *fiber.Ctx) error {
	userEmail := new(validators.VerifyOTP)

	if err := c.BodyParser(userEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
}
