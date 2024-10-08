package expenses

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-pdf/fpdf"
	"github.com/spf13/cast"
	"github.com/tech/core/domain"
	"math"
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

	response.TotalYearExpenses = cast.ToString(totalYearExpenses)

	userAllDetails, err := t.expensePersistence.ExpenseDetailsPersistence.GetYpExpenseDateById(user.Id, mapIdAndMonth[month], year)
	if err != nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "452", Msg: err.Error()}
	}

	if userAllDetails == nil {
		fmt.Println(funcName, "no Expenses of user")
		return domain.Response{Code: "200", Msg: "expenses not found", Model: response}
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
		return domain.Response{Code: "200", Msg: "expenses not found", Model: response}
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

	//userDetailsByPayment, err := t.expensePersistence.ExpenseDetailsPersistence.GetUserExpenseByUidAndPaymentMode(user.Id, paymentMode, mapIdAndMonth[month], year)
	//if err != nil {
	//	fmt.Println(funcName, "no Expenses of user err :", err)
	//	return domain.Response{Code: "452", Msg: "err.Error()"}
	//}
	//
	//if userDetailsByPayment == nil {
	//	fmt.Println(funcName, "no Expenses of user")
	//	return domain.Response{Code: "453", Msg: "expenses not found"}
	//}

	//response.PaymentModeNames = paymentModes
	//response.UserPaymentExpense = userDetailsByPayment

	return domain.Response{Code: "200", Msg: "success", Model: response}
}

func (t *Expenses) GetMonthExpensesPdf(ctx context.Context, user *domain.User, month int, year int) domain.Response {
	funcName := "GetMonthExpensesPdf"
	fmt.Println(funcName)

	addCategoryNames := make(map[string]float64)
	addPaymentModes := make(map[string]float64)
	var totalYearExpenses float64
	var totalMonthExpenses float64

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
			totalMonthExpenses = cast.ToFloat64(details.ExpensesAmount)
		}
	}

	userAllDetails, err := t.expensePersistence.ExpenseDetailsPersistence.GetYpExpenseDateById(user.Id, mapIdAndMonth[month], year)
	if err != nil {
		fmt.Println(funcName, "no Expenses of user err :", err)
		return domain.Response{Code: "452", Msg: err.Error()}
	}

	if userAllDetails == nil {
		fmt.Println(funcName, "no Expenses of user")
		return domain.Response{Code: "200", Msg: "expenses not found"}
	}

	for _, details := range userAllDetails {
		addCategoryNames[details.Category] += cast.ToFloat64(details.Amount)
		addPaymentModes[details.PaymentMode] += cast.ToFloat64(details.Amount)
	}

	pdfBuffer, err := t.GeneratePDFForMonthExpenses(user, mapIdAndMonth[month], cast.ToString(year), totalYearExpenses, totalMonthExpenses, addCategoryNames, addPaymentModes)
	if err != nil {
		fmt.Println(funcName, "error while generate PDF :", err)
	}

	subject := ""
	body := ""
	err = t.message.SendEmailWithPDF(pdfBuffer, user.Email, subject, body)
	if err != nil {
		fmt.Println(funcName, "error while send email :", err)
	}

	return domain.Response{Code: "200", Msg: "success"}
}

func (t *Expenses) GeneratePDFForMonthExpenses(user *domain.User, month string, year string, totalYearExpenses float64, totalMonthExpenses float64, addCategoryNames map[string]float64, addPaymentModes map[string]float64) ([]byte, error) {

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(5, 5, 5)
	pdf.AddPage()

	pageWidth, _ := pdf.GetPageSize()
	// Title
	pdf.SetFont("Arial", "B", 23)
	title := "Expenses of " + month + " " + year
	titleWidth := pdf.GetStringWidth(title)
	pdf.SetX((pageWidth - titleWidth) / 2)
	pdf.Cell(0, 40, title)
	pdf.Ln(30)

	pdf.SetFont("Arial", "", 19)
	totalYear := "Total amount spends " + year + " := " + cast.ToString(totalYearExpenses)
	subtitleWidth := pdf.GetStringWidth(totalYear)
	pdf.SetX((pageWidth - subtitleWidth) / 2)
	pdf.Cell(0, 10, totalYear)
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 19)
	totalMonth := "Total amount spends " + month + " := " + cast.ToString(totalMonthExpenses)
	totalMonthWidth := pdf.GetStringWidth(totalMonth)
	pdf.SetX((pageWidth - totalMonthWidth) / 2)
	pdf.Cell(0, 10, totalMonth)
	pdf.Ln(18)

	pdf.SetFont("Arial", "B", 13)
	pdf.SetTextColor(68, 114, 196)
	categoryText := "Spends Category"
	categoryTextWidth := pdf.GetStringWidth(categoryText)
	pdf.SetX((pageWidth - categoryTextWidth) / 2)
	pdf.Cell(0, 10, categoryText)
	pdf.Ln(10)
	pdf.SetTextColor(0, 0, 0)

	// Table Header
	pdf.SetFont("Arial", "B", 14)
	pdf.SetFillColor(200, 220, 255) // Light blue background
	headers := []string{"Category", "Amount", "Percentage"}
	colWidths := []float64{40, 40, 40}

	totalTableWidth := 0.0
	for _, width := range colWidths {
		totalTableWidth += width
	}

	startX := (pageWidth - totalTableWidth) / 2
	pdf.SetX(startX)

	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 10, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	counts := 0
	// Table Rows
	pdf.SetFont("Arial", "", 12)
	for category, amount := range addCategoryNames {
		if counts%2 == 0 {
			pdf.SetFillColor(240, 240, 240) // Light gray for even rows
		} else {
			pdf.SetFillColor(255, 255, 255) // White for odd rows
		}
		percent := (amount / totalMonthExpenses) * 100
		roundedPercent := math.Round(percent)
		rowData := []string{
			category,
			fmt.Sprintf("%.2f", amount),
			fmt.Sprintf("%.2f", roundedPercent),
		}

		pdf.SetX(startX) // Align rows to startX for centered table
		for j, data := range rowData {
			pdf.CellFormat(colWidths[j], 10, data, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)
		counts++
	}

	pdf.SetFont("Arial", "B", 13)
	pdf.SetTextColor(68, 114, 196)
	categoryText = "Spends PaymentMode"
	categoryTextWidth = pdf.GetStringWidth(categoryText)
	pdf.SetX((pageWidth - categoryTextWidth) / 2)
	pdf.Cell(0, 15, categoryText)
	pdf.Ln(12)
	pdf.SetTextColor(0, 0, 0)

	// Table Header
	pdf.SetFont("Arial", "B", 14)
	pdf.SetFillColor(200, 220, 255) // Light blue background
	headers = []string{"PaymentMode", "Amount", "Percentage"}

	pdf.SetX(startX)

	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 10, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	counts = 0
	// Table Rows
	pdf.SetFont("Arial", "", 12)
	for paymentName, amount := range addPaymentModes {
		if counts%2 == 0 {
			pdf.SetFillColor(240, 240, 240) // Light gray for even rows
		} else {
			pdf.SetFillColor(255, 255, 255) // White for odd rows
		}
		percent := (amount / totalMonthExpenses) * 100
		roundedPercent := math.Round(percent)
		rowData := []string{
			paymentName,
			fmt.Sprintf("%.2f", amount),
			fmt.Sprintf("%.2f", roundedPercent),
		}

		pdf.SetX(startX) // Align rows to startX for centered table
		for j, data := range rowData {
			pdf.CellFormat(colWidths[j], 10, data, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)
	}

	// Output to byte array
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
