package driver

import (
	"context"
	"github.com/tech/core/domain"
)

type LoginService interface {
	GenerateOtpForUser(c context.Context, otpRequest *domain.OtpRequest) domain.Response
}
