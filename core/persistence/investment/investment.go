package investment

import (
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/internal/beego"
)

type Investment struct {
	StockNamePersistence     YpStockName
	InvestmentBuyPersistence YpInvestmentBuy
}

func InvestmentPersistence() *Investment {
	return &Investment{
		StockNamePersistence:     YpStockNameDetails(),
		InvestmentBuyPersistence: YpInvestmentBuyDetails(),
	}
}

type YpStockName interface {
	GetAllStockNames() (data []*domain.StockNameInJson, err error)
}

func YpStockNameDetails() YpStockName {
	return &beego.BeegoStockName{}
}

type YpInvestmentBuy interface {
	AddYpInvestmentBuyDetails(newEntry *domain.InvestmentBuyDetails) (int64, error)
}

func YpInvestmentBuyDetails() YpInvestmentBuy {
	return &beego.BeegoInvestmentBuyDetails{}
}
