package domain

import "time"

type ExpenseDate struct {
	Id        int
	Uid       int
	Date      int
	Month     string
	Year      int
	Amount    string
	UpdatedOn time.Time
	CreatedOn time.Time
}

type MonthIncomeExpense struct {
	Id             int
	Uid            int
	Month          string
	Year           int
	IncomeAmount   string
	ExpensesAmount string
	UpdatedOn      time.Time
	CreatedOn      time.Time
}
