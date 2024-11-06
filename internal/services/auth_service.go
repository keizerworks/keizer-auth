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

func (as *AuthService) RegisterUser(userRegister *validators.UserRegister) error {
	passwordHash, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	otp, err := utils.GenerateOTP()
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	err = SendOTPEmail(userRegister.Email, otp)
	if err != nil {
		return fmt.Errorf("failed to send OTP email: %w", err)
	}

	err = as.userRepo.CreateUser(&models.User{
		Email:        userRegister.Email,
		FirstName:    userRegister.FirstName,
		LastName:     userRegister.LastName,
		PasswordHash: passwordHash,
		Otp:          otp,
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
