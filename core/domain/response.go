package domain

type Response struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

type OtpResponse struct {
	OtpDigits   int    `json:"otpDigits"`
	Token       string `json:"token"`
	CustomToken string `json:"customToken"`
	UserState   string `json:"userState"`
}

type UserExpenseYearResponse struct {
	IsOldUser bool     `json:"isOldUser"`
	Labels    []string `json:"labels"`
	Expenses  []string `json:"expenses"`
}
