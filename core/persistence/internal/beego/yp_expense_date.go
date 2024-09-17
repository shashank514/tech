package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoExpenseDate struct{}

func (t *BeegoExpenseDate) AddYpExpenseDate(data *domain.ExpenseDate) (id int64, err error) {
	v := t.convertDomainToModel(data)
	id, err = v.AddYpExpenseDate()
	return
}

func (t *BeegoExpenseDate) GetYpExpenseDateById(uid int, month string, year int) (data []*domain.ExpenseDate, err error) {
	v, err := new(ypmodel.YpExpenseDate).GetYpExpenseDateById(uid, month, year)
	if err == nil {
		for _, details := range v {
			data = append(data, t.convertModelToDomain(details))
		}
		return data, nil
	}
	return nil, err
}

func (t *BeegoExpenseDate) convertDomainToModel(data *domain.ExpenseDate) *ypmodel.YpExpenseDate {
	return &ypmodel.YpExpenseDate{
		Id:        data.Id,
		Uid:       data.Uid,
		Date:      data.Date,
		Month:     data.Month,
		Year:      data.Year,
		Amount:    data.Amount,
		CreatedOn: data.CreatedOn,
		UpdatedOn: data.UpdatedOn,
	}
}

func (t *BeegoExpenseDate) convertModelToDomain(data *ypmodel.YpExpenseDate) *domain.ExpenseDate {
	return &domain.ExpenseDate{
		Id:        data.Id,
		Uid:       data.Uid,
		Date:      data.Date,
		Month:     data.Month,
		Year:      data.Year,
		Amount:    data.Amount,
		CreatedOn: data.CreatedOn,
		UpdatedOn: data.UpdatedOn,
	}
}
