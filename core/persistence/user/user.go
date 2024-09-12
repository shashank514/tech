package user

import (
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/internal/beego"
)

type User struct {
	YpUserPersistence  YpUser
	UserOtpPersistence YpUserOtp
}

func NewUser() *User {
	return &User{
		YpUserPersistence:  YpUserDetails(),
		UserOtpPersistence: YpUserOtpDetails(),
	}
}

type YpUser interface {
	AddYPUser(user *domain.User) (id int64, err error)
	GetYPUserByEmail(email string) (user *domain.User, err error)
}

func YpUserDetails() YpUser {
	return &beego.BeegoYpUser{}
}

type YpUserOtp interface {
	AddYUserOtp(otp *domain.UserOtp) (id int64, err error)
	GetYpUserOtpByToken(token string) (otp *domain.UserOtp, err error)
	UpdateYpUserOtpByColumn(data *domain.UserOtp, column ...string) (err error)
}

func YpUserOtpDetails() YpUserOtp {
	return &beego.BeegoUserOtp{}
}
