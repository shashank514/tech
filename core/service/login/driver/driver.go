package driver

import (
	"context"
	"github.com/tech/core/domain"
)

type LoginService interface {
	GenerateOtpForUser(c context.Context, otpRequest *domain.OtpRequest) domain.Response
	SubmitOtpForUser(c context.Context, otpRequest *domain.OtpRequest) domain.Response
	GetUserPersonalDetails(exeCtx context.Context, user *domain.User, personalDetails *domain.PersonalDetails) domain.Response
	GetUserAddressDetails(exeCtx context.Context, user *domain.User, userAddress *domain.AddressDetails) domain.Response
	GetUserDetails(exeCtx context.Context, user *domain.User) domain.Response
	GetNewToken(requestBody *domain.NewToken) domain.Response
}
