package investment

import (
	"context"
	"fmt"
	"github.com/tech/core/domain"
)

var mapIdAndMonth = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func (s *Investment) AddNewInvestmentDetails(ctx context.Context, user *domain.User, request *domain.InvestmentBuyRequest) domain.Response {
	funcName := "AddNewInvestmentDetails"
	newEntry := &domain.InvestmentBuyDetails{
		Uid:            user.Id,
		Date:           request.Date,
		Month:          mapIdAndMonth[request.Month],
		Year:           request.Year,
		Type:           request.Category,
		Name:           request.Name,
		Symbol:         request.Symbol,
		Enable:         1,
		BuyCount:       request.BuyCount,
		AmountPerBuy:   request.AmountPerBuy,
		TotalAmount:    request.TotalAmount,
		RemainingCount: request.BuyCount,
		FdInterest:     request.RateOfInterest,
	}

	// insert to yp_investment_buy_details table
	_, err := s.investment.InvestmentBuyPersistence.AddYpInvestmentBuyDetails(newEntry)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: err.Error()}
	}

	return domain.Response{Code: "200", Msg: "success"}
}
