package validators

type VerifyOTP struct {
	Otp string `validate:"required" label:"OTP"`
	Id  string `validate:"required" label:"Id"`
}
