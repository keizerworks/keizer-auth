package controllers

import (
	"errors"
	"fmt"
	"keizer-auth/internal/models"
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
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid request body"})
	}

	user, err := ac.authService.GetUser(body.Email)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Unable to retrieve user information. Please try again later.",
			})
	}
	if user.ID.String() == "" {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "User not found. Please check your email and try again.",
			})
	}
	if !user.IsVerified {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "User is not verified. Please verify your account before signing in.",
			})
	}
	if user.Type != models.Dashboard {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	isValid, err := ac.authService.VerifyPassword(
		body.Password,
		user.PasswordHash,
	)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Unable to verify password. Please try again later."})
	}
	if !isValid {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid email or password. Please try again."})
	}

	sessionId, err := ac.sessionService.CreateSession(user)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Something went wrong, Failed to create session"})
	}

	fmt.Printf("%v", sessionId)
	fmt.Print(sessionId)
	utils.SetSessionCookie(c, sessionId)
	return c.JSON(user)
}

func (ac *AuthController) SignUp(c *fiber.Ctx) error {
	user := new(validators.SignUpUser)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if valid, errors := user.Validate(); !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	id, err := ac.authService.RegisterUser(user)
	if err != nil {
		fmt.Print("&+v\n", err)
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

	return c.JSON(fiber.Map{"id": id, "message": "User Signed Up!"})
}

func (ac *AuthController) VerifyOTP(c *fiber.Ctx) error {
	verifyOtpBody := new(validators.VerifyOTP)

	if err := c.BodyParser(verifyOtpBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	userID, isOtpValid, err := ac.authService.VerifyOTP(verifyOtpBody)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "otp expired" {
			return c.
				Status(fiber.StatusNotFound).
				JSON(fiber.Map{"error": "OTP expired"})
		}

		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to verify OTP"})
	}

	if !isOtpValid {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "OTP not valid"})
	}

	user, err := ac.authService.SetIsVerified(userID)
	if err != nil {
		return err
	}

	sessionID, err := ac.sessionService.CreateSession(user)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to create session"})
	}

	utils.SetSessionCookie(c, sessionID)
	return c.JSON(user)
}

func (ac *AuthController) Profile(c *fiber.Ctx) error {
	user := utils.GetCurrentUser(c)
	return c.JSON(user)
}
