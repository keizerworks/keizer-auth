package validators

import "github.com/gookit/validate"

type UserRegister struct {
	Email           string `json:"email" validate:"required|email" label:"Email"`
	Password        string `json:"password" validate:"required|minLen:8|regex:^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[!@#~$%^&*(),.?\":{}|<>]).+$" label:"Password"`
	ConfirmPassword string `json:"confirm_password" validate:"required|minLen:8|eqField:password" label:"Confirm Password"`
	FirstName       string `json:"first_name" validate:"required|maxLen:32" label:"First Name"`
	LastName        string `json:"last_name" validate:"maxLen:32" label:"Last Name"`
}

func (f UserRegister) Messages() map[string]string {
	return validate.MS{
		// Global messages
		"required":                "{field} is required",
		"email":                   "Please provide a valid email address",
		"minLen":                  "{field} must be at least 8 characters",
		"maxLen":                  "{field} cannot exceed 32 characters",
		"ConfirmPassword.eqField": "{field} does not match Password",

		// Field-specific messages
		"Password.regex": "Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character",
	}
}

func (u *UserRegister) Validate() (bool, map[string]map[string]string) {
	v := validate.Struct(u)

	if !v.Validate() {
		return false, v.Errors.All()
	}

	return true, nil
}
