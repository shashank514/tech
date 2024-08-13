package login

type LoginResponse struct {
	Otp  string `json:"otp"`
	Name string `json:"name"`
}

func GenerateOtp() *LoginResponse {
	return &LoginResponse{
		Otp:  "1234555",
		Name: "shashank k p",
	}
}
