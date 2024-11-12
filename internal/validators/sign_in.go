package validators

type SignInUser struct {
	Email    string `validate:"required|email" label:"Email"`
	Password string `validate:"required|minLen:8|password" label:"Password"`
}
