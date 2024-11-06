package services

import (
	"errors"
	"keizer-auth-api/internal/utils"
	"net/smtp"
	"os"
)

var (
	from = os.Getenv("EMAIL_FROM")
)

type EmailService struct {
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (es *EmailService) SendEmail(to string, message string) error {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "local" {
		// Send email using mailhog
		sendEmailUsingMailHog(to, message)
	} else if appEnv == "production" {
		// Send email using SendGrid
	} else {
		return errors.New("invalid APP_ENV")
	}
	return nil
}

func sendEmailUsingMailHog(to string, message string) error {
	smtpHost := "localhost"
	smtpPort := "1025"
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func SendOTPEmail(to string, otp string) error {
	emailService := NewEmailService()
	message := utils.ConstructOTPMail("OTP Verification", "Your OTP is "+otp)
	err := emailService.SendEmail(to, message)
	if err != nil {
		return err
	}
	return nil
}
