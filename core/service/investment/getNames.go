package investment

import (
	"context"
	"fmt"
	"github.com/tech/cloud/message"
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/investment"
	"github.com/tech/core/service/investment/driver"
	"github.com/tech/utilities"
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

var bankName = []string{
	"HDFC Bank",
	"State Bank of India",
	"ICICI Bank",
	"Punjab National Bank",
	"Axis Bank",
	"Kotak Mahindra Bank",
	"Bank of Baroda",
	"Canara Bank",
	"IndusInd Bank",
	"IDBI Bank",
	"Bank of India",
	"Union Bank of India",
	"Central Bank of India",
	"Indian Bank",
	"Indian Overseas Bank",
	"UCO Bank",
	"Yes Bank",
	"IDFC FIRST Bank",
	"Federal Bank",
	"RBL Bank",
	"South Indian Bank",
	"Bandhan Bank",
	"Jammu & Kashmir Bank",
	"Punjab & Sind Bank",
	"Karur Vysya Bank",
	"Dhanlaxmi Bank",
	"City Union Bank",
	"Tamilnad Mercantile Bank",
	"Suryoday Small Finance Bank",
	"AU Small Finance Bank",
	"Ujjivan Small Finance Bank",
	"Equitas Small Finance Bank",
	"Fincare Small Finance Bank",
}

func (s *Investment) GetSelectNames(ctx context.Context, user *domain.User, category string) domain.Response {
	funcName := "GetSelectNames"
	var response []*domain.StockNameInJson
	var err error

	switch category {
	case utilities.ConstStock, utilities.ConstMutualFund:
		response, err = s.investment.StockNamePersistence.GetAllYpDetailsNameByCategory(category)
		if err != nil {
			fmt.Println(funcName, " error while getting stock name err : ", err)
		}
	case utilities.ConstFd:
		for _, name := range bankName {
			response = append(response, &domain.StockNameInJson{Name: name})
		}
	}
	return domain.Response{Code: "200", Msg: "success", Model: response}
}
