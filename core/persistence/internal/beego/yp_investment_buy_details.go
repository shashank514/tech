package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoInvestmentBuyDetails struct{}

func (t *BeegoInvestmentBuyDetails) AddYpInvestmentBuyDetails(newEntry *domain.InvestmentBuyDetails) (int64, error) {
	v := t.convertDomainToModel(newEntry)
	return v.AddYpInvestmentBuyDetails()
}

func (t *BeegoInvestmentBuyDetails) convertDomainToModel(data *domain.InvestmentBuyDetails) *ypmodel.YpInvestmentBuyDetails {
	return &ypmodel.YpInvestmentBuyDetails{
		Id:             data.Id,
		Uid:            data.Uid,
		Date:           data.Date,
		Month:          data.Month,
		Year:           data.Year,
		Type:           data.Type,
		Name:           data.Name,
		Symbol:         data.Symbol,
		Enable:         data.Enable,
		BuyCount:       data.BuyCount,
		AmountPerBuy:   data.AmountPerBuy,
		TotalAmount:    data.TotalAmount,
		RemainingCount: data.RemainingCount,
		FdInterest:     data.FdInterest,
		CreatedOn:      data.CreatedOn,
		UpdatedOn:      data.UpdatedOn,
	}
}

func (t *BeegoInvestmentBuyDetails) convertModelToDomain(data *ypmodel.YpInvestmentBuyDetails) *domain.InvestmentBuyDetails {
	return &domain.InvestmentBuyDetails{
		Id:             data.Id,
		Uid:            data.Uid,
		Date:           data.Date,
		Month:          data.Month,
		Year:           data.Year,
		Type:           data.Type,
		Name:           data.Name,
		Symbol:         data.Symbol,
		Enable:         data.Enable,
		BuyCount:       data.BuyCount,
		AmountPerBuy:   data.AmountPerBuy,
		TotalAmount:    data.TotalAmount,
		RemainingCount: data.RemainingCount,
		FdInterest:     data.FdInterest,
		CreatedOn:      data.CreatedOn,
		UpdatedOn:      data.UpdatedOn,
	}
}
