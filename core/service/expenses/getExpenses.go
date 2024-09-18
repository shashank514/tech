package expenses

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tech/core/domain"
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

	fmt.Println("user id ", user.Id, "month ", month, "year ", year)

	userAllDetails, err := t.expensePersistence.ExpenseDetailsPersistence.GetYpExpenseDateById(user.Id, mapIdAndMonth[month], year)
	if err != nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "452", Msg: "err.Error()"}
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
		return domain.Response{Code: "452", Msg: "err.Error()"}
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

	//for _, details := range userDetailsByDate {
	//	DateLabels = append(DateLabels, cast.ToString(details.Date))
	//	DateExpenses = append(DateExpenses, cast.ToString(details.Amount))
	//}
	response.DateLabels = DateLabels
	response.DateExpenses = DateExpenses

	return domain.Response{Code: "200", Msg: "success", Model: response}
}
