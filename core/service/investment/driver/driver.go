package driver

import (
	"context"
	"github.com/tech/core/domain"
)

type InvestmentService interface {
	GetSelectNames(ctx context.Context, user *domain.User, category string) domain.Response
}
