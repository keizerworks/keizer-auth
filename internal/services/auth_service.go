package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"keizer-auth/internal/models"
	"keizer-auth/internal/repositories"
	"keizer-auth/internal/utils"
	"keizer-auth/internal/validators"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	userRepo  *repositories.UserRepository
	redisRepo *repositories.RedisRepository
}

func NewAuthService(userRepo *repositories.UserRepository, redisRepo *repositories.RedisRepository) *AuthService {
	return &AuthService{userRepo: userRepo, redisRepo: redisRepo}
}

func (as *AuthService) RegisterUser(
	userRegister *validators.SignUpUser,
) (string, error) {
	user := models.User{
		Email: userRegister.Email,
	}

	fmt.Print(user)
	fmt.Print(user.IsVerified)
	err := as.userRepo.GetUserByStruct(&user)
	if err != nil {
		return "", err
	}

	if user.IsVerified {
		return "", fmt.Errorf("uesr already exists")
	}

	passwordHash, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	user.FirstName = userRegister.FirstName
	user.LastName = userRegister.LastName
	user.PasswordHash = passwordHash

	otp, err := utils.GenerateOTP()
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}

	otpCacheKey := cuid2.Generate()
	hashOtp, err := utils.HashPassword(otp)
	if err != nil {
		return "", fmt.Errorf("failed to hash OTP: %w", err)
	}

	err = as.userRepo.CreateUser(&user)
	if err != nil {
		return "", err
	}

	otpData := models.OTPData{
		OTPHash: hashOtp,
		ID:      user.ID.String(),
	}

	marshalledOtpData, err := json.Marshal(otpData)
	if err != nil {
		return "", err
	}

	encodedOtpData := base64.StdEncoding.EncodeToString(marshalledOtpData)
	err = as.redisRepo.Set(otpCacheKey, encodedOtpData, time.Minute*3)
	if err != nil {
		return "", fmt.Errorf("failed to save otp in redis: %w", err)
	}

	// TODO: track status, add reties
	go SendOTPEmail(userRegister.Email, otp)

	return otpCacheKey, nil
}

func (as *AuthService) VerifyPassword(
	password string,
	passwordHash string,
) (bool, error) {
	return utils.VerifyPassword(password, passwordHash)
}

func (as *AuthService) VerifyOTP(verifyOtpBody *validators.VerifyOTP) (string, bool, error) {
	encodedOtpData, err := as.redisRepo.Get(verifyOtpBody.Id)
	if err != nil {
		if err == redis.Nil {
			return "", false, fmt.Errorf("otp expired")
		}
		return "", false, fmt.Errorf("failed to get otp from redis %w", err)
	}

	decodedOtpData, err := base64.StdEncoding.DecodeString(encodedOtpData)
	if err != nil {
		return "", false, fmt.Errorf("failed to decode otp data: %w", err)
	}

	var otpData models.OTPData
	err = json.Unmarshal(decodedOtpData, &otpData)
	if err != nil {
		return "", false, fmt.Errorf("failed to unmarshal otp data: %w", err)
	}

	isVerified, err := utils.VerifyPassword(verifyOtpBody.Otp, otpData.OTPHash)
	return otpData.ID, isVerified, err
}

func (as *AuthService) SetIsVerified(id string) (*models.User, error) {
	user := models.User{IsVerified: true}
	err := as.userRepo.UpdateUser(id, &user)
	return &user, err
}

func (as *AuthService) GetUser(email string) (*models.User, error) {
	user := models.User{Email: email}
	err := as.userRepo.GetUserByStruct(&user)
	return &user, err
}
