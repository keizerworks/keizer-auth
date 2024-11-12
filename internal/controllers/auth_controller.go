package controllers

import (
	"errors"

	"keizer-auth/internal/services"
	"keizer-auth/internal/utils"
	"keizer-auth/internal/validators"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	authService    *services.AuthService
	sessionService *services.SessionService
}

func NewAuthController(as *services.AuthService, ss *services.SessionService) *AuthController {
	return &AuthController{authService: as, sessionService: ss}
}

func (ac *AuthController) SignIn(c *fiber.Ctx) error {
	body := new(validators.SignInUser)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	isValid, userDetails, err := ac.authService.VerifyPassword(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	sessionId, err := ac.sessionService.CreateSession(userDetails.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create session"})
	}
	utils.SetSessionCookie(c, sessionId)

	return c.JSON(fiber.Map{"message": "signed in successfully"})
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
	verifyOtpBody := new(validators.VerifyOTP)

	if err := c.BodyParser(verifyOtpBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	isOtpValid, err := ac.authService.VerifyOTP(verifyOtpBody)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "OTP not found"})
		}
		if err.Error() == "otp expired" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "OTP expired"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to verify OTP"})
	}
	if !isOtpValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "OTP not valid"})
	}

	err = ac.authService.SetIsVerified(verifyOtpBody.Id)
	if err != nil {
		return err
	}

	sessionId, err := ac.sessionService.CreateSession(verifyOtpBody.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create session"})
	}
	utils.SetSessionCookie(c, sessionId)

	return c.JSON(fiber.Map{"message": "OTP Verified!"})
}
