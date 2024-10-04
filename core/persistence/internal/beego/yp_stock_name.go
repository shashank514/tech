package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoStockName struct{}

func (t *BeegoStockName) GetAllStockNames() (data []*domain.StockNameInJson, err error) {
	v, err := new(ypmodel.YpStockName).GetAllYpStockName()
	if err == nil {
		for _, details := range v {
			data = append(data, t.ConvertToStockNameInJson(details))
		}
		return data, nil
	}
	return nil, err
}

func (t *BeegoStockName) GetYpStockNameBySymbol(symbol string) (data *domain.StockName, err error) {
	v, err := new(ypmodel.YpStockName).GetYpStockNameBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	return t.convertModelToDomain(v), nil
}

func (t *BeegoStockName) convertDomainToModel(data *domain.StockName) *ypmodel.YpStockName {
	return &ypmodel.YpStockName{
		Id:        data.Id,
		StockName: data.StockName,
		Symbol:    data.Symbol,
		Price:     data.Price,
		UpdatedOn: data.UpdatedOn,
		CreatedOn: data.CreatedOn,
	}
}

func (t *BeegoStockName) convertModelToDomain(data *ypmodel.YpStockName) *domain.StockName {
	return &domain.StockName{
		Id:        data.Id,
		StockName: data.StockName,
		Symbol:    data.Symbol,
		Price:     data.Price,
		UpdatedOn: data.UpdatedOn,
		CreatedOn: data.CreatedOn,
	}
}

func (t *BeegoStockName) ConvertToStockNameInJson(data *ypmodel.YpStockName) *domain.StockNameInJson {
	return &domain.StockNameInJson{
		Name:   data.StockName,
		Symbol: data.Symbol,
	}
}
