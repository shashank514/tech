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

type InvestmentBuyRequest struct {
	Date           int    `json:"date" binding:"required"`
	Month          int    `json:"month" binding:"required"`
	Year           int    `json:"year" binding:"required"`
	Category       string `json:"category" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Symbol         string `json:"symbol"`
	BuyCount       string `json:"buyCount"`
	AmountPerBuy   string `json:"amountPerBuy"`
	TotalAmount    string `json:"totalAmount" binding:"required"`
	RateOfInterest string `json:"rateOfInterest"`
}
