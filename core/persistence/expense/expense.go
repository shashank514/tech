package expense

import (
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/internal/beego"
)

type Expenses struct {
	ExpenseDatePersistence        YpExpenseDate
	MonthIncomeExpensePersistence YpMonthExpenseIncome
}

func ExpensesDetails() *Expenses {
	return &Expenses{
		ExpenseDatePersistence:        YpExpenseDateDetails(),
		MonthIncomeExpensePersistence: YpMonthExpenseIncomeDetails(),
	}
}

type YpExpenseDate interface {
	AddYpExpenseDate(data *domain.ExpenseDate) (id int64, err error)
	GetYpExpenseDateById(uid int, month string, year int) (data []*domain.ExpenseDate, err error)
}

func YpExpenseDateDetails() YpExpenseDate {
	return &beego.BeegoExpenseDate{}
}

type YpMonthExpenseIncome interface {
	AddMonthIncomeExpense(data *domain.MonthIncomeExpense) (id int64, err error)
	GetMonthIncomeExpenseByYear(uid int, year int) (data []*domain.MonthIncomeExpense, err error)
}

func YpMonthExpenseIncomeDetails() YpMonthExpenseIncome {
	return &beego.BeegoMonthIncomeExpense{}
}
