package investment

import (
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/internal/beego"
)

type Investment struct {
	StockNamePersistence YpStockName
}

func InvestmentPersistence() *Investment {
	return &Investment{
		StockNamePersistence: YpStockNameDetails(),
	}
}

type YpStockName interface {
	GetAllStockNames() (data []*domain.StockNameInJson, err error)
}

func YpStockNameDetails() YpStockName {
	return &beego.BeegoStockName{}
}
