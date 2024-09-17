package expenses

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/tech/cloud/message"
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/expense"
	"github.com/tech/core/service/expenses/driver"
	"time"
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

var mapIdAndMonth = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

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

func (t *Expenses) AddExpense(ctx context.Context, user *domain.User, request *domain.ExpenseRequest) domain.Response {
	funcName := "AddExpense"

	newEnter := &domain.Expense{
		Uid:         user.Id,
		Date:        request.Date,
		Month:       mapIdAndMonth[request.Month],
		Year:        request.Year,
		Amount:      request.Amount,
		Category:    request.Category,
		PaymentMode: request.PaymentMode,
		Description: request.Description,
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	}

	// add user expenses
	_, err := t.expensePersistence.ExpenseDetailsPersistence.AddUserExpense(newEnter)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: "error while adding new entry"}
	}

	// get data from yp_expense_date
	dateDetails, err := t.expensePersistence.ExpenseDatePersistence.GetYpExpenseDateByUidAndDate(user.Id, request.Date, mapIdAndMonth[request.Month], request.Year)
	if err == orm.ErrNoRows {
		newDateEntry := &domain.ExpenseDate{
			Uid:       user.Id,
			Date:      request.Date,
			Month:     mapIdAndMonth[request.Month],
			Year:      request.Year,
			Amount:    request.Amount,
			CreatedOn: time.Now(),
			UpdatedOn: time.Now(),
		}

		_, err = t.expensePersistence.ExpenseDatePersistence.AddYpExpenseDate(newDateEntry)
		if err != nil {
			fmt.Println(funcName, err)
			return domain.Response{Code: "452", Msg: "error while adding new entry"}
		}
	} else if err == nil {

		// add the amount to existing amount
		dateDetails.Amount += request.Amount
		dateDetails.UpdatedOn = time.Now()
		updateColumns := []string{
			"amount",
			"updatedOn",
		}

		err = t.expensePersistence.ExpenseDatePersistence.UpdateYpExpenseDateByColumns(dateDetails, updateColumns...)
		if err != nil {
			fmt.Println(funcName, err)
			return domain.Response{Code: "452", Msg: "error while Updating entry"}
		}

	} else if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "453", Msg: "error while getting details"}
	}

	//
	monthDetails, err := t.expensePersistence.MonthIncomeExpensePersistence.GetDetailsUsingUidAndMonth(user.Id, mapIdAndMonth[request.Month], request.Year)
	if err == orm.ErrNoRows {

		monthNewEntry := &domain.MonthIncomeExpense{
			Uid:            user.Id,
			Month:          mapIdAndMonth[request.Month],
			Year:           request.Year,
			ExpensesAmount: request.Amount,
			CreatedOn:      time.Now(),
			UpdatedOn:      time.Now(),
		}

		_, err = t.expensePersistence.MonthIncomeExpensePersistence.AddMonthIncomeExpense(monthNewEntry)
		if err != nil {
			fmt.Println(funcName, err)
			return domain.Response{Code: "452", Msg: "error while adding new entry"}
		}
	} else if err == nil {

		monthDetails.ExpensesAmount += request.Amount
		monthDetails.UpdatedOn = time.Now()
		updateColumns := []string{
			"income_amount",
			"updatedOn",
		}

		err = t.expensePersistence.MonthIncomeExpensePersistence.UpdateMonthIncomeExpenseByColumns(monthDetails, updateColumns...)
		if err != nil {
			fmt.Println(funcName, err)
			return domain.Response{Code: "452", Msg: "error while Updating entry"}
		}
	} else if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: "error while getting details"}
	}

	return domain.Response{Code: "200", Msg: "success"}
}
