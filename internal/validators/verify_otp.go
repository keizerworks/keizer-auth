package validators

type VerifyOTP struct {
	Email string `validate:"required|email" label:"Email"`
	Otp   string `validate:"required" label:"OTP"`
	Id    string `validate:"required" label:"Id"`
}
