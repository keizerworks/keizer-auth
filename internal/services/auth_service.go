package services

import (
	"fmt"
	"keizer-auth-api/internal/models"
	"keizer-auth-api/internal/repositories"
	"keizer-auth-api/internal/utils"
	"keizer-auth-api/internal/validators"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (as *AuthService) RegisterUser(userRegister *validators.SignUpUser) error {
	passwordHash, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	// TODO: email should be sent using async func
	if err = SendOTPEmail(userRegister.Email, otp); err != nil {
		return fmt.Errorf("failed to send OTP email: %w", err)
	}

	if err = as.userRepo.CreateUser(&models.User{
		Email:        userRegister.Email,
		FirstName:    userRegister.FirstName,
		LastName:     userRegister.LastName,
		PasswordHash: passwordHash,
	}); err != nil {
		return err
	}

	return nil
}
