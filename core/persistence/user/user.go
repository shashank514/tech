package user

import (
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/internal/beego"
	"time"
)

type User struct {
	YpUserPersistence      YpUser
	UserOtpPersistence     YpUserOtp
	UserAddressPersistence YpUserAddress
}

func NewUser() *User {
	return &User{
		YpUserPersistence:      YpUserDetails(),
		UserOtpPersistence:     YpUserOtpDetails(),
		UserAddressPersistence: YpUserAddressDetails(),
	}
}

type YpUser interface {
	AddYPUser(user *domain.User) (id int64, err error)
	GetYPUserByEmail(email string) (user *domain.User, err error)
	GetYpUserByAuth(auth string) (user *domain.User, err error)
	UpdateYpUserByColumn(data *domain.User, column ...string) (err error)
}

func YpUserDetails() YpUser {
	return &beego.BeegoYpUser{}
}

type YpUserOtp interface {
	AddYUserOtp(otp *domain.UserOtp) (id int64, err error)
	GetYpUserOtpByToken(token string) (otp *domain.UserOtp, err error)
	UpdateYpUserOtpByColumn(data *domain.UserOtp, column ...string) (err error)
	GetYpUserOtpCount(sentTo string, today time.Time) (counts int, data *domain.UserOtp)
}

func YpUserOtpDetails() YpUserOtp {
	return &beego.BeegoUserOtp{}
}

type YpUserAddress interface {
	AddYpUserAddress(address *domain.UserAddress) (id int64, err error)
}

func YpUserAddressDetails() YpUserAddress {
	return &beego.BeegoUserAddress{}
}
