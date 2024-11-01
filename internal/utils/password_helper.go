package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	Memory      = 64 * 1024 // 64 MB
	Iterations  = 3         // Number of iterations
	Parallelism = 2         // Number of threads
	SaltLength  = 16        // Salt length in bytes
	KeyLength   = 32        // Desired length of the hash
)

func GenerateRandomString() ([]byte, error) {
	bytes := make([]byte, 12)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}

func GenerateHash(value []byte, salt []byte) []byte {
	return argon2.IDKey([]byte(value), salt, Iterations, Memory, Parallelism, KeyLength)
}

func HashPassword(password string) (string, error) {
	salt, err := GenerateRandomString()
	if err != nil {
		return "", err
	}

	hash := GenerateHash([]byte(password), salt)
	encodedHash := fmt.Sprintf(
		"%s.%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encodedHash, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, ".")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	computedHash := GenerateHash([]byte(password), salt)
	return subtle.ConstantTimeCompare(hash, computedHash) == 1, nil
}
