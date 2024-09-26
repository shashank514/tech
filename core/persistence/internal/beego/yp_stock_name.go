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

func (t *BeegoStockName) ConvertToStockNameInJson(data *ypmodel.YpStockName) *domain.StockNameInJson {
	return &domain.StockNameInJson{
		Name:   data.StockName,
		Symbol: data.Symbol,
	}
}
