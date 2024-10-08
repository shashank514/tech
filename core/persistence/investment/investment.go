package investment

import (
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/internal/beego"
)

type Investment struct {
	StockNamePersistence         YpStockName
	InvestmentBuyPersistence     YpInvestmentBuy
	InvestmentDetailsPersistence YpInvestmentDetails
}

func InvestmentPersistence() *Investment {
	return &Investment{
		StockNamePersistence:         YpStockNameDetails(),
		InvestmentBuyPersistence:     YpInvestmentBuyDetails(),
		InvestmentDetailsPersistence: InvestmentDetails(),
	}
}

type YpStockName interface {
	GetAllStockNames() (data []*domain.StockNameInJson, err error)
	GetYpStockNameBySymbol(symbol string) (data *domain.StockName, err error)
	GetAllYpDetailsNameByCategory(category string) (data []*domain.StockNameInJson, err error)
}

func YpStockNameDetails() YpStockName {
	return &beego.BeegoStockName{}
}

type YpInvestmentBuy interface {
	AddYpInvestmentBuyDetails(newEntry *domain.InvestmentBuyDetails) (int64, error)
	GetAllYpInvestmentBuyDetailsByUid(uid int) (data []*domain.InvestmentBuyDetails, err error)
	GetInvestmentBuyDetailsByType(types string, uid int) (data []*domain.InvestmentBuyDetails, err error)
}

func YpInvestmentBuyDetails() YpInvestmentBuy {
	return &beego.BeegoInvestmentBuyDetails{}
}

type YpInvestmentDetails interface {
	AddYpInvestmentDetails(data *domain.InvestmentDetails) (int64, error)
	GetUserInvestmentDetailsByUid(uid int) (*domain.InvestmentDetails, error)
	UpdateYpInvestmentDetailsByColumns(data *domain.InvestmentDetails, column ...string) (err error)
}

func InvestmentDetails() YpInvestmentDetails {
	return &beego.BeegoInvestmentDetails{}
}
