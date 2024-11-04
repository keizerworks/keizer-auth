package services

import (
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
	// TODO: create and send OTP to user email for user verification

	passwordHash, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return err
	}

	return as.userRepo.CreateUser(&models.User{
		Email:        userRegister.Email,
		FirstName:    userRegister.FirstName,
		LastName:     userRegister.LastName,
		PasswordHash: passwordHash,
	})
}
