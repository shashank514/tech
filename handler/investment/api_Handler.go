package investment

import (
	"github.com/gin-gonic/gin"
	"github.com/tech/cloud/message"
	"github.com/tech/core/persistence/investment"
	"github.com/tech/core/persistence/user"
	investment2 "github.com/tech/core/service/investment"
	investment3 "github.com/tech/handler/investment/investment"
	"github.com/tech/middleware"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.Use(middleware.DBConnectionWithEnvMiddleware)
	router.Use(middleware.CustomerAuthMiddleware(user.YpUserDetails()))

	messages := message.NewMessage()
	investmentPersist := *investment.InvestmentPersistence()

	investmentSvc := investment2.AllInvestmentService(messages, investmentPersist)

	getName := investment3.GetSelectCategoryDetails(investmentSvc)
	addNewInvestment := investment3.UserNewInvestment(investmentSvc)
	investmentDetails := investment3.GetUserAllInvestments(investmentSvc)
	getHoldings := investment3.GetInvestedHoldings(investmentSvc)

	router.GET("/getName", getName)
	router.POST("/addInvestment", addNewInvestment)
	router.GET("/getInvestmentDetails", investmentDetails)
	router.GET("/holdings", getHoldings)
}
