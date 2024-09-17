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
		year := cast.ToInt(c.Query("year"))

		user := c.Keys["customer"]
		if user == nil {
			fmt.Println(functionName, "invalid credentials")
			response := domain.Response{Code: "459", Msg: "Session has expired"}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		loggedUser := user.(*domain.User)

		response := svc.GetUserYearExpenses(exeCtx, loggedUser, year)
		c.JSON(http.StatusOK, response)
	}
}
