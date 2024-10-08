package userExpense

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/tech/core/domain"
	"github.com/tech/core/service/expenses/driver"
	"github.com/tech/util"
	"net/http"
)

func GetUserExpense(svc driver.ExpenseService) gin.HandlerFunc {
	functionName := "GetUserExpense"
	var exeCtx context.Context

	return func(c *gin.Context) {
		exeCtx = util.SetContext(c)

		types := cast.ToString(c.Query("type"))

		user := c.Keys["customer"]
		if user == nil {
			fmt.Println(functionName, "invalid credentials")
			response := domain.Response{Code: "459", Msg: "Session has expired"}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		loggedUser := user.(*domain.User)

		response := domain.Response{}
		switch types {
		case "home":
			year := cast.ToInt(c.Query("year"))
			response = svc.GetUserYearExpenses(exeCtx, loggedUser, year)
		case "expenses":
			month := cast.ToInt(c.Query("month"))
			year := cast.ToInt(c.Query("year"))
			response = svc.GetUserExpenses(exeCtx, loggedUser, month, year)
		case "transaction":
			category := cast.ToString(c.Query("category"))
			paymentMode := cast.ToString(c.Query("paymentMode"))
			month := cast.ToInt(c.Query("month"))
			year := cast.ToInt(c.Query("year"))
			response = svc.GetUserExpensesByCategorys(exeCtx, loggedUser, category, paymentMode, month, year)
		default:
			response.Code = "400"
			response.Msg = "invalid Query par"
		}
		c.JSON(http.StatusOK, response)
	}
}

func AddUserExpense(svc driver.ExpenseService) gin.HandlerFunc {
	functionName := "AddUserExpense"
	var exeCtx context.Context
	return func(c *gin.Context) {
		exeCtx = util.SetContext(c)

		user := c.Keys["customer"]
		if user == nil {
			fmt.Println(functionName, "invalid credentials")
			response := domain.Response{Code: "459", Msg: "Session has expired"}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		requestBody := &domain.ExpenseRequest{}
		if err := c.Bind(requestBody); err != nil {
			fmt.Println(err)
			response := domain.Response{Code: "404", Msg: "invalid request body"}
			c.JSON(http.StatusOK, response)
			return
		}

		loggedUser := user.(*domain.User)

		response := svc.AddExpense(exeCtx, loggedUser, requestBody)
		c.JSON(http.StatusOK, response)
	}
}

func GetMonthExpensesPdf(svc driver.ExpenseService) gin.HandlerFunc {
	functionName := "GetMonthExpensesPdf"
	var exeCtx context.Context
	return func(c *gin.Context) {
		exeCtx = util.SetContext(c)

		user := c.Keys["customer"]
		if user == nil {
			fmt.Println(functionName, "invalid credentials")
			response := domain.Response{Code: "459", Msg: "Session has expired"}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		loggedUser := user.(*domain.User)

		month := cast.ToInt(c.Query("month"))
		year := cast.ToInt(c.Query("year"))

		response := svc.GetMonthExpensesPdf(exeCtx, loggedUser, month, year)
		c.JSON(http.StatusOK, response)
	}
}
