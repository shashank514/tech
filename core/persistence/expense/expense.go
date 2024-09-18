package expense

import (
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/internal/beego"
)

type Expenses struct {
	ExpenseDatePersistence        YpExpenseDate
	MonthIncomeExpensePersistence YpMonthExpenseIncome
	ExpenseDetailsPersistence     YpExpenseDetails
}

func ExpensesDetails() *Expenses {
	return &Expenses{
		ExpenseDatePersistence:        YpExpenseDateDetails(),
		MonthIncomeExpensePersistence: YpMonthExpenseIncomeDetails(),
		ExpenseDetailsPersistence:     YpExpenseDetailsDetails(),
	}
}

type YpExpenseDate interface {
	AddYpExpenseDate(data *domain.ExpenseDate) (id int64, err error)
	GetYpExpenseDateById(uid int, month string, year int) (data []*domain.ExpenseDate, err error)
	GetYpExpenseDateByUidAndDate(uid int, date int, month string, year int) (data *domain.ExpenseDate, err error)
	UpdateYpExpenseDateByColumns(date *domain.ExpenseDate, columns ...string) error
}

func YpExpenseDateDetails() YpExpenseDate {
	return &beego.BeegoExpenseDate{}
}

type YpMonthExpenseIncome interface {
	AddMonthIncomeExpense(data *domain.MonthIncomeExpense) (id int64, err error)
	GetMonthIncomeExpenseByYear(uid int, year int) (data []*domain.MonthIncomeExpense, err error)
	GetDetailsUsingUidAndMonth(uid int, month string, year int) (data *domain.MonthIncomeExpense, err error)
	UpdateMonthIncomeExpenseByColumns(date *domain.MonthIncomeExpense, columns ...string) error
}

func YpMonthExpenseIncomeDetails() YpMonthExpenseIncome {
	return &beego.BeegoMonthIncomeExpense{}
}

type YpExpenseDetails interface {
	AddUserExpense(details *domain.Expense) (id int64, err error)
	GetYpExpenseDateById(uid int, month string, year int) (data []*domain.Expense, err error)
}

func YpExpenseDetailsDetails() YpExpenseDetails {
	return &beego.BeegoExpenseDetails{}
}
