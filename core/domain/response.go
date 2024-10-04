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
	TotalYearExpenses  string   `json:"totalYearExpenses"`
	TotalMonthExpenses string   `json:"totalMonthExpenses"`
}

type CategoryExpenseResponse struct {
	CategoryNames       []string   `json:"categoryNames"`
	UserCategoryExpense []*Expense `json:"categoryExpense"`
	PaymentModeNames    []string   `json:"paymentModeNames"`
	UserPaymentExpense  []*Expense `json:"paymentExpense"`
}

type InvestmentDetailsResponse struct {
	UserInvested                  bool                       `json:"userInvested"`
	TotalInvestedAmount           string                     `json:"totalInvestedAmount"`
	PresentInvestedAmount         string                     `json:"presentInvestedAmount"`
	ProfitAfter                   string                     `json:"profitAfter"`
	LossAfter                     string                     `json:"lossAfter"`
	InvestmentCategoryLabels      []string                   `json:"investmentCategoryLabels"`
	InvestmentCategoriesExpenses  []string                   `json:"investmentCategoriesExpenses"`
	InvestmentCategoriesTable     []InvestmentDetailsToTable `json:"investmentCategoriesTable"`
	StockInvestmentLabels         []string                   `json:"stockInvestmentLabels"`
	StockInvestmentExpense        []string                   `json:"stockInvestmentExpense"`
	StockInvestmentTable          []InvestmentDetailsToTable `json:"stockInvestmentTable"`
	MutualFundsInvestmentLabels   []string                   `json:"mutualFundsInvestmentLabels"`
	MutualFundsInvestmentExpenses []string                   `json:"mutualFundsInvestmentExpenses"`
	MutualFundsInvestmentTable    []InvestmentDetailsToTable `json:"mutualFundsInvestmentTable"`
	FDInvestmentLabels            []string                   `json:"fdInvestmentLabels"`
	FDInvestmentExpense           []string                   `json:"fdInvestmentExpense"`
	FDInvestmentTable             []InvestmentDetailsToTable `json:"fdInvestmentTable"`
}

type InvestmentDetailsToTable struct {
	Name       string `json:"name"`
	Amount     string `json:"amount"`
	Percentage string `json:"percentage"`
}

type InvestmentHoldingResponse struct {
	StockInvestment  string           `json:"stockInvestment"`
	StockCurrent     string           `json:"stockCurrent"`
	StockTotalReturn string           `json:"stockTotalReturn"`
	AllShareDetails  []HoldingDetails `json:"allShareDetails"`
}

type HoldingDetails struct {
	Name           string `json:"name"`
	Quantity       string `json:"quantity"`
	AvgPrice       string `json:"avgPrice"`
	MktPrice       string `json:"mktPrice"`
	InvestedAmount string `json:"investedAmount"`
	CurrentAmount  string `json:"currentAmount"`
	TotalReturns   string `json:"totalReturns"`
}
