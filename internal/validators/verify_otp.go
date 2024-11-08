package validators

type VerifyOTP struct {
	Email string `validate:"required|email" label:"Email"`
}
