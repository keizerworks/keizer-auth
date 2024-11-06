package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	otpLength = 6
)

func GenerateOTP() (string, error) {
	otp := ""
	for i := 0; i < otpLength; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return "", err
		}
		otp += n.String()
	}
	return otp, nil
}
