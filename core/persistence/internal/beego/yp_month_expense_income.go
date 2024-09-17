package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoMonthIncomeExpense struct {
}

func (t *BeegoMonthIncomeExpense) AddMonthIncomeExpense(data *domain.MonthIncomeExpense) (id int64, err error) {
	v := t.convertDomainToModel(data)
	return v.AddYpMonthIncomeExpense()
}

func (t *BeegoMonthIncomeExpense) GetMonthIncomeExpenseByYear(uid int, year int) (data []*domain.MonthIncomeExpense, err error) {
	v, err := new(ypmodel.YpMonthIncomeExpense).GetYpMonthIncomeExpenseByYear(uid, year)
	if err == nil {
		for _, details := range v {
			data = append(data, t.convertModelToDomain(details))
		}
		return data, nil
	}
	return nil, err
}

func (t *BeegoMonthIncomeExpense) GetDetailsUsingUidAndMonth(uid int, month string, year int) (data *domain.MonthIncomeExpense, err error) {
	v, err := new(ypmodel.YpMonthIncomeExpense).GetDetailsUsingUidAndMonth(uid, month, year)
	if err != nil {
		return nil, err
	}
	return t.convertModelToDomain(v), nil
}

func (t *BeegoMonthIncomeExpense) UpdateMonthIncomeExpenseByColumns(date *domain.MonthIncomeExpense, columns ...string) error {
	v := t.convertDomainToModel(date)
	return v.UpdateYpMonthIncomeExpenseByColumn(columns...)
}

func (t *BeegoMonthIncomeExpense) convertDomainToModel(data *domain.MonthIncomeExpense) *ypmodel.YpMonthIncomeExpense {
	return &ypmodel.YpMonthIncomeExpense{
		Id:             data.Id,
		Uid:            data.Uid,
		Month:          data.Month,
		Year:           data.Year,
		IncomeAmount:   data.IncomeAmount,
		ExpensesAmount: data.ExpensesAmount,
		CreatedOn:      data.CreatedOn,
		UpdatedOn:      data.UpdatedOn,
	}
}

func (t *BeegoMonthIncomeExpense) convertModelToDomain(data *ypmodel.YpMonthIncomeExpense) *domain.MonthIncomeExpense {
	return &domain.MonthIncomeExpense{
		Id:             data.Id,
		Uid:            data.Uid,
		Month:          data.Month,
		Year:           data.Year,
		IncomeAmount:   data.IncomeAmount,
		ExpensesAmount: data.ExpensesAmount,
		CreatedOn:      data.CreatedOn,
		UpdatedOn:      data.UpdatedOn,
	}
}
