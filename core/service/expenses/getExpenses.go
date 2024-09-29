package expenses

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tech/core/domain"
	"strings"
)

func (t *Expenses) GetUserExpenses(ctx context.Context, user *domain.User, month int, year int) domain.Response {
	funcName := "GetUserExpenses"
	response := domain.UserExpenseResponse{}
	addCategoryNames := make(map[string]int)
	addDateAndAmount := make(map[int]int)
	var CategoryLabels []string
	var CategoriesExpenses []string
	var DateLabels []string
	var DateExpenses []string
	var totalYearExpenses float64

	fmt.Println("user id ", user.Id, "month ", month, "year ", year)

	userAllDetails, err := t.expensePersistence.ExpenseDetailsPersistence.GetYpExpenseDateById(user.Id, mapIdAndMonth[month], year)
	if err != nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "452", Msg: err.Error()}
	}

	if userAllDetails == nil {

		fmt.Println(funcName, "no Expenses of user")
		return domain.Response{Code: "453", Msg: "expenses not found"}
	}

	for _, details := range userAllDetails {
		addCategoryNames[details.Category] += cast.ToInt(details.Amount)
	}

	for category, amount := range addCategoryNames {
		CategoryLabels = append(CategoryLabels, category)
		CategoriesExpenses = append(CategoriesExpenses, cast.ToString(amount))
	}
	response.CategoryLabels = CategoryLabels
	response.CategoriesExpenses = CategoriesExpenses

	userDetailsByDate, err := t.expensePersistence.ExpenseDatePersistence.GetYpExpenseDateById(user.Id, mapIdAndMonth[month], year)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: err.Error()}
	}

	if userDetailsByDate == nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "453", Msg: "expenses not found"}
	}

	for _, details := range userDetailsByDate {
		addDateAndAmount[details.Date] = cast.ToInt(details.Amount)
	}

	for i := 1; i <= 31; i++ {
		DateLabels = append(DateLabels, cast.ToString(i))
		DateExpenses = append(DateExpenses, cast.ToString(addDateAndAmount[i]))
	}

	response.DateLabels = DateLabels
	response.DateExpenses = DateExpenses

	userMonthExpenses, err := t.expensePersistence.MonthIncomeExpensePersistence.GetMonthIncomeExpenseByYear(user.Id, year)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: err.Error()}
	}

	if userMonthExpenses == nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "453", Msg: "expenses not found"}
	}

	for _, details := range userMonthExpenses {
		totalYearExpenses += cast.ToFloat64(details.ExpensesAmount)
		if details.Month == mapIdAndMonth[month] {
			response.TotalMonthExpenses = cast.ToString(details.ExpensesAmount)
		}
	}

	response.TotalMonthExpenses = cast.ToString(totalYearExpenses)

	return domain.Response{Code: "200", Msg: "success", Model: response}
}

func (t *Expenses) GetUserExpensesByCategorys(ctx context.Context, user *domain.User, categorys string, paymentMode string, month int, year int) domain.Response {
	funcName := "GetUserExpensesByCategorys"
	response := domain.CategoryExpenseResponse{}

	var userAllDetails []*domain.Expense
	var err error

	if categorys == "all" {
		userAllDetails, err = t.expensePersistence.ExpenseDetailsPersistence.GetYpExpenseDateById(user.Id, mapIdAndMonth[month], year)
	} else {
		userAllDetails, err = t.expensePersistence.ExpenseDetailsPersistence.GetUserExpenseByUidAndCategory(user.Id, categorys, mapIdAndMonth[month], year)
	}

	if err != nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "452", Msg: "err.Error()"}
	}

	if userAllDetails == nil {
		fmt.Println(funcName, "no Expenses of user")
		return domain.Response{Code: "453", Msg: "expenses not found"}
	}

	monthDetails, err := t.expensePersistence.MonthIncomeExpensePersistence.GetDetailsUsingUidAndMonth(user.Id, mapIdAndMonth[month], year)
	if err != nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "452", Msg: "err.Error()"}
	}

	response.CategoryNames = strings.Split(monthDetails.ExpensesCategory, ",")
	response.CategoryNames = append(response.CategoryNames, "all")
	response.UserCategoryExpense = userAllDetails

	userDetailsByPayment, err := t.expensePersistence.ExpenseDetailsPersistence.GetUserExpenseByUidAndPaymentMode(user.Id, paymentMode, mapIdAndMonth[month], year)
	if err != nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "452", Msg: "err.Error()"}
	}

	if userDetailsByPayment == nil {
		fmt.Println(funcName, "no Expenses of user")
		return domain.Response{Code: "453", Msg: "expenses not found"}
	}

	//response.PaymentModeNames = paymentModes
	response.UserPaymentExpense = userDetailsByPayment

	return domain.Response{Code: "200", Msg: "success", Model: response}
}
