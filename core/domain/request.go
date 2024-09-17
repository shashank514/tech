package domain

type OtpRequest struct {
	Action string `json:"action" binding:"required,alpha"`
	Email  string `json:"email"`
	Otp    string `json:"otp"`
	Token  string `json:"token"`
}
