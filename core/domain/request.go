package domain

type OtpRequest struct {
	Action string `json:"action" binding:"required,alpha"`
	Email  string `json:"email"`
	Otp    string `json:"otp"`
	Token  string `json:"token"`
}

type ExpenseRequest struct {
	Date        int    `json:"date" binding:"required"`
	Month       int    `json:"month" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
	Category    string `json:"category" binding:"required"`
	PaymentMode string `json:"paymentMode" binding:"required"`
	Description string `json:"description"`
}
