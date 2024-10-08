package domain

import "time"

type StockName struct {
	Id        int
	Enable    int
	StockName string
	Symbol    string
	Price     float64
	Category  string
	Sector    string
	Industry  string
	UpdatedOn time.Time
	CreatedOn time.Time
}

type StockNameInJson struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type InvestmentBuyDetails struct {
	Id             int
	Uid            int
	Date           int
	Month          string
	Year           int
	Type           string
	Name           string
	Symbol         string
	Enable         int
	BuyCount       string
	AmountPerBuy   string
	TotalAmount    string
	RemainingCount string
	FdInterest     string
	UpdatedOn      time.Time
	CreatedOn      time.Time
}

type InvestmentDetails struct {
	Id                      int
	Uid                     int
	TotalInvestmentAmount   string
	PresentInvestmentAmount string
	ProfitAfterSellAmount   string
	LossAfterSellAmount     string
	UpdatedOn               time.Time
	CreatedOn               time.Time
}
