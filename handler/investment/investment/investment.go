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

//func UserNewInvestment()

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
