package expense

import (
	"github.com/gin-gonic/gin"
	"github.com/tech/cloud/message"
	"github.com/tech/core/persistence/expense"
	"github.com/tech/core/persistence/user"
	"github.com/tech/core/service/expenses"
	"github.com/tech/handler/expense/userExpense"
	"github.com/tech/middleware"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.Use(middleware.DBConnectionWithEnvMiddleware)
	router.Use(middleware.CustomerAuthMiddleware(user.YpUserDetails()))

	messages := message.NewMessage()
	expensePersist := *expense.ExpensesDetails()

	expenseSvc := expenses.ExpensesDetails(messages, expensePersist)

	getUserExpense := userExpense.GetUserExpense(expenseSvc)

	router.GET("/getExpenses", getUserExpense)
}
