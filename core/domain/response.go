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

type UserExpenseResponse struct {
	CategoryLabels     []string `json:"categoryLabels"`
	CategoriesExpenses []string `json:"categoriesExpenses"`
	DateLabels         []string `json:"dateLabels"`
	DateExpenses       []string `json:"dateExpenses"`
}

type CategoryExpenseResponse struct {
	CategoryNames       []string   `json:"categoryNames"`
	UserCategoryExpense []*Expense `json:"categoryExpense"`
	PaymentModeNames    []string   `json:"paymentModeNames"`
	UserPaymentExpense  []*Expense `json:"paymentExpense"`
}
