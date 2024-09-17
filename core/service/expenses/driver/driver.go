package driver

import (
	"context"
	"github.com/tech/core/domain"
)

type ExpenseService interface {
	GetUserYearExpenses(ctx context.Context, user *domain.User, year int) domain.Response
}
