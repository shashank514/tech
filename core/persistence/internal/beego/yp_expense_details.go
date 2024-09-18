package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoExpenseDetails struct{}

func (t *BeegoExpenseDetails) AddUserExpense(details *domain.Expense) (id int64, err error) {
	v := t.convertDomainToModel(details)
	return v.AddUserExpense()
}

func (t *BeegoExpenseDetails) GetYpExpenseDateById(uid int, month string, year int) (data []*domain.Expense, err error) {
	v, err := new(ypmodel.YpExpenseDetails).GetUserExpenseById(uid, month, year)
	if err == nil {
		for _, details := range v {
			data = append(data, t.convertModelToDomain(details))
		}
		return data, nil
	}
	return nil, err
}

func (t *BeegoExpenseDetails) convertDomainToModel(data *domain.Expense) *ypmodel.YpExpenseDetails {
	return &ypmodel.YpExpenseDetails{
		Id:          data.Id,
		Uid:         data.Uid,
		Date:        data.Date,
		Month:       data.Month,
		Year:        data.Year,
		Amount:      data.Amount,
		Category:    data.Category,
		PaymentMode: data.PaymentMode,
		Description: data.Description,
		CreatedOn:   data.CreatedOn,
		UpdatedOn:   data.UpdatedOn,
	}
}

func (t *BeegoExpenseDetails) convertModelToDomain(data *ypmodel.YpExpenseDetails) *domain.Expense {
	return &domain.Expense{
		Id:          data.Id,
		Uid:         data.Uid,
		Date:        data.Date,
		Month:       data.Month,
		Year:        data.Year,
		Amount:      data.Amount,
		Category:    data.Category,
		PaymentMode: data.PaymentMode,
		Description: data.Description,
		CreatedOn:   data.CreatedOn,
		UpdatedOn:   data.UpdatedOn,
	}
}
