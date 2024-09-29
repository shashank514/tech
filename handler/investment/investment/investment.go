package investment

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/tech/core/domain"
	"github.com/tech/core/service/investment/driver"
	"github.com/tech/util"
	"net/http"
)

func GetSelectCategoryDetails(svc driver.InvestmentService) gin.HandlerFunc {
	functionName := "GetSelectCategoryDetails"
	var exeCtx context.Context

	return func(c *gin.Context) {
		exeCtx = util.SetContext(c)

		category := cast.ToString(c.Query("category"))

		user := c.Keys["customer"]
		if user == nil {
			fmt.Println(functionName, "invalid credentials")
			response := domain.Response{Code: "459", Msg: "Session has expired"}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		loggedUser := user.(*domain.User)

		response := svc.GetSelectNames(exeCtx, loggedUser, category)

		c.JSON(http.StatusOK, response)
	}
}

func UserNewInvestment(svc driver.InvestmentService) gin.HandlerFunc {
	functionName := "UserNewInvestment"
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

		requestBody := &domain.InvestmentBuyRequest{}
		if err := c.Bind(requestBody); err != nil {
			fmt.Println(err)
			response := domain.Response{Code: "404", Msg: "invalid request body"}
			c.JSON(http.StatusOK, response)
			return
		}

		loggedUser := user.(*domain.User)

		response := svc.AddNewInvestmentDetails(exeCtx, loggedUser, requestBody)
		c.JSON(http.StatusOK, response)
	}
}

func GetUserAllInvestments(svc driver.InvestmentService) gin.HandlerFunc {
	functionName := "GetUserAllInvestments"
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

		response := svc.GetInvestmentDetails(exeCtx, loggedUser)

		c.JSON(http.StatusOK, response)
	}
}
