package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoInvestmentDetails struct{}

func (t *BeegoInvestmentDetails) AddYpInvestmentDetails(data *domain.InvestmentDetails) (int64, error) {
	v := t.convertDomainToModel(data)
	return v.AddYpInvestmentDetails()
}

func (t *BeegoInvestmentDetails) GetUserInvestmentDetailsByUid(uid int) (*domain.InvestmentDetails, error) {
	v, err := new(ypmodel.YPInvestmentDetails).GetUserInvestmentDetailsByUid(uid)
	if err != nil {
		return nil, err
	}
	return t.convertModelToDomain(v), nil
}

func (t *BeegoInvestmentDetails) UpdateYpInvestmentDetailsByColumns(data *domain.InvestmentDetails, column ...string) (err error) {
	v := t.convertDomainToModel(data)
	return v.UpdateYpInvestmentDetailsByColumns(column...)
}

func (t *BeegoInvestmentDetails) convertDomainToModel(data *domain.InvestmentDetails) *ypmodel.YPInvestmentDetails {
	return &ypmodel.YPInvestmentDetails{
		Id:                      data.Id,
		Uid:                     data.Uid,
		TotalInvestmentAmount:   data.TotalInvestmentAmount,
		PresentInvestmentAmount: data.PresentInvestmentAmount,
		ProfitAfterSellAmount:   data.ProfitAfterSellAmount,
		LossAfterSellAmount:     data.LossAfterSellAmount,
		UpdatedOn:               data.UpdatedOn,
		CreatedOn:               data.CreatedOn,
	}
}

func (t *BeegoInvestmentDetails) convertModelToDomain(data *ypmodel.YPInvestmentDetails) *domain.InvestmentDetails {
	return &domain.InvestmentDetails{
		Id:                      data.Id,
		Uid:                     data.Uid,
		TotalInvestmentAmount:   data.TotalInvestmentAmount,
		PresentInvestmentAmount: data.PresentInvestmentAmount,
		ProfitAfterSellAmount:   data.ProfitAfterSellAmount,
		LossAfterSellAmount:     data.LossAfterSellAmount,
		UpdatedOn:               data.UpdatedOn,
		CreatedOn:               data.CreatedOn,
	}
}
