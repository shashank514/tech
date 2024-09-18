package driver

import (
	"context"
	"github.com/tech/core/domain"
)

type ExpenseService interface {
	GetUserYearExpenses(ctx context.Context, user *domain.User, year int) domain.Response
	AddExpense(ctx context.Context, user *domain.User, request *domain.ExpenseRequest) domain.Response
	GetUserExpenses(ctx context.Context, user *domain.User, month int, year int) domain.Response
}
