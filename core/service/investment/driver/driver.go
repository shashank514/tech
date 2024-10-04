package driver

import (
	"context"
	"github.com/tech/core/domain"
)

type InvestmentService interface {
	GetSelectNames(ctx context.Context, user *domain.User, category string) domain.Response
	AddNewInvestmentDetails(ctx context.Context, user *domain.User, request *domain.InvestmentBuyRequest) domain.Response
	GetInvestmentDetails(ctx context.Context, user *domain.User) domain.Response
	GetUserHoldings(ctx context.Context, user *domain.User) domain.Response
}
