package domain

import "time"

type User struct {
	Id        int
	Auth      string
	Email     string
	Mobile    string
	Name      string
	CreatedOn time.Time
	UpdatedOn time.Time
}

type UserOtp struct {
	Id        int
	Uid       int
	SentOn    time.Time
	Validity  time.Time
	Otp       string
	Token     string
	SentTo    string
	Tries     int
	Validated int8
	UpdatedOn time.Time
	CreatedOn time.Time
}

type Response struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

type OtpRequest struct {
	Action string `json:"action" binding:"required,alpha"`
	Email  string `json:"email"`
	Otp    string `json:"otp"`
	Token  string `json:"token"`
}

type OtpResponse struct {
	OtpDigits   int    `json:"otpDigits"`
	Token       string `json:"token"`
	CustomToken string `json:"customToken"`
	UserState   string `json:"userState"`
}
