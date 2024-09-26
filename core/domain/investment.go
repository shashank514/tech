package domain

import "time"

type StockName struct {
	Id        int
	StockName string
	Symbol    string
	Price     float64
	UpdatedOn time.Time
	CreatedOn time.Time
}

type StockNameInJson struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
