package validators

import (
	"unicode"

	"keizer-auth/internal/utils"

	"github.com/gookit/validate"
)

func init() {
	validate.AddValidator("password", func(password string) bool {
		var hasDigit, hasUpperLetter, hasLowerLetter, hasPunct bool

		for _, r := range password {
			if unicode.IsDigit(r) && !hasDigit {
				hasDigit = true
			}
			if unicode.IsUpper(r) && !hasUpperLetter {
				hasUpperLetter = true
			}
			if unicode.IsLower(r) && !hasLowerLetter {
				hasLowerLetter = true
			}
			if unicode.IsPunct(r) && !hasPunct {
				hasPunct = true
			}
		}

		return hasDigit && hasUpperLetter && hasLowerLetter && hasPunct
	})
}

type SignUpUser struct {
	Email     string `validate:"required|email" label:"Email"`
	Password  string `validate:"required|minLen:8|password" label:"Password"`
	FirstName string `json:"first_name" validate:"required|maxLen:32" label:"First Name"`
	LastName  string `json:"last_name" validate:"maxLen:32" label:"Last Name"`
}

func (f SignUpUser) Messages() map[string]string {
	return validate.MS{
		// Global messages
		"required": "{field} is required",
		"email":    "Please provide a valid email address",
		"minLen":   "{field} must be at least 8 characters",
		"maxLen":   "{field} cannot exceed 32 characters",
		"password": "Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character",
	}
}

func (u *SignUpUser) Validate() (bool, map[string]map[string]string) {
	v := validate.Struct(u)

	if !v.Validate() {
		// Convert error field names to snake_case
		errors := v.Errors.All()
		snakeCaseErrors := make(map[string]map[string]string)

		for field, fieldErrors := range errors {
			snakeCaseField := utils.ToSnakeCase(field)
			snakeCaseErrors[snakeCaseField] = fieldErrors
		}

		return false, snakeCaseErrors
	}

	return true, nil
}
