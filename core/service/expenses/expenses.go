package expenses

import (
	"context"
	"fmt"
	"github.com/tech/cloud/message"
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/expense"
	"github.com/tech/core/service/expenses/driver"
)

type Expenses struct {
	message            message.MessageDriver
	expensePersistence expense.Expenses
}

func ExpensesDetails(message message.MessageDriver, expensePersistence expense.Expenses) driver.ExpenseService {
	return &Expenses{
		message:            message,
		expensePersistence: expensePersistence,
	}
}

var monthNames = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func (t *Expenses) GetUserYearExpenses(ctx context.Context, user *domain.User, year int) domain.Response {
	funcName := "GetUserYearExpenses"
	response := domain.UserExpenseYearResponse{}
	var labels []string
	var expenses []string

	// get user expenses of the year from yp_month_expense_income
	expenseDetails, err := t.expensePersistence.MonthIncomeExpensePersistence.GetMonthIncomeExpenseByYear(user.Id, year)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: "err.Error()"}
	}

	if expenseDetails == nil {
		return domain.Response{Code: "453", Msg: "expenses not found"}
	}

	expenseInGroup := make(map[string]*domain.MonthIncomeExpense, len(expenseDetails))
	checkMonthDetails := make(map[string]bool, len(expenseDetails))

	for _, expenseDetail := range expenseDetails {
		expenseInGroup[expenseDetail.Month] = expenseDetail
		checkMonthDetails[expenseDetail.Month] = true
	}

	for _, month := range monthNames {
		if checkMonthDetails[month] {
			labels = append(labels, month)
			expenses = append(expenses, expenseInGroup[month].ExpensesAmount)
		} else {
			labels = append(labels, month)
			expenses = append(expenses, "0")
		}
	}

	response.Labels = labels
	response.Expenses = expenses

	return domain.Response{Code: "200", Msg: "success", Model: response}
}
