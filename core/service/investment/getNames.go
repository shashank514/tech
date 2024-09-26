package investment

import (
	"context"
	"fmt"
	"github.com/tech/cloud/message"
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/investment"
	"github.com/tech/core/service/investment/driver"
)

type Investment struct {
	message    message.MessageDriver
	investment investment.Investment
}

func AllInvestmentService(message message.MessageDriver, investment investment.Investment) driver.InvestmentService {
	return &Investment{
		message:    message,
		investment: investment,
	}
}

func (s *Investment) GetSelectNames(ctx context.Context, user *domain.User, category string) domain.Response {
	funcName := "GetSelectNames"
	var response []*domain.StockNameInJson
	var err error

	switch category {
	case "Stock":
		response, err = s.investment.StockNamePersistence.GetAllStockNames()
		if err != nil {
			fmt.Println(funcName, " error while getting stock name err : ", err)
		}
	}
	return domain.Response{Code: "200", Msg: "success", Model: response}
}
