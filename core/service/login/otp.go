package login

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/tech/cloud/message"
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/user"
	"github.com/tech/core/service/login/driver"
	"time"
)

type Login struct {
	message         message.MessageDriver
	userPersistence user.User
}

func NewLoginDetails(message message.MessageDriver, userPersistence user.User) driver.LoginService {
	return &Login{
		message:         message,
		userPersistence: userPersistence,
	}
}

func (b *Login) GenerateOtpForUser(c context.Context, otpRequest *domain.OtpRequest) domain.Response {

	function := "GenerateOtpForUser"
	response := &domain.OtpResponse{}
	otpDigits := 6

	// get user details from yp_user
	userDetails, err := b.userPersistence.YpUserPersistence.GetYPUserByEmail(otpRequest.Email)
	if err != nil && err != orm.ErrNoRows {
		fmt.Println(function, err)
		return domain.Response{Code: "452", Msg: "err while user details"}
	} else if err == orm.ErrNoRows {
		userDetails = &domain.User{
			Email: otpRequest.Email,
		}

		_, err = b.userPersistence.YpUserPersistence.AddYPUser(userDetails)
		if err != nil {
			fmt.Println(function, err)
			return domain.Response{Code: "452", Msg: "err while creating user"}
		}
	}

	response.OtpDigits = otpDigits
	otp := GenerateRandomOtp(otpDigits)
	token := RandAlphanumeric(12)
	response.Token = token

	err = b.AddUserOtpToDB(userDetails.Id, otp, otpRequest.Email, token)
	if err != nil {
		fmt.Println(function, err)
		return domain.Response{Code: "452", Msg: "err while insert user new otp in DB"}
	}

	subject := "Subject: Your OTP Code\r\n"
	body := fmt.Sprintf("Your OTP code is: %s", otp)

	err = b.message.SendEmail(otpRequest.Email, subject, body)
	if err != nil {
		fmt.Println(function, err)
		return domain.Response{Code: "452", Msg: "err while send email"}
	}

	return domain.Response{Code: "200", Msg: "success", Model: response}
}

func (b *Login) SubmitOtpForUser(c context.Context, otpRequest *domain.OtpRequest) domain.Response {
	function := "SubmitOtpForUser"
	response := &domain.OtpResponse{}
	response.OtpDigits = 6
	var updateColumns []string

	// get user details from yp_user
	userDetails, err := b.userPersistence.YpUserPersistence.GetYPUserByEmail(otpRequest.Email)
	if err != nil {
		fmt.Println(function, err)
		return domain.Response{Code: "452", Msg: "err while user details"}
	}

	// get user otp from yp_user_otp table using token
	otpDetails, err := b.userPersistence.UserOtpPersistence.GetYpUserOtpByToken(otpRequest.Token)
	if err != nil {
		fmt.Println(function, err)
		return domain.Response{Code: "452", Msg: "err while get otp from token"}
	}

	if otpDetails.Tries > 3 {
		fmt.Println(function, "Submit invalid otp 3 times")
		return domain.Response{Code: "452", Msg: "Submit invalid otp 3 times"}
	}

	if otpDetails.SentTo != otpRequest.Email {
		fmt.Println(function, "email is invalid")
		return domain.Response{Code: "452", Msg: "email is invalid"}
	}

	otpDetails.Tries += 1
	updateColumns = append(updateColumns, "tries")

	if otpDetails.Otp != otpRequest.Otp {
		err = b.userPersistence.UserOtpPersistence.UpdateYpUserOtpByColumn(otpDetails, updateColumns...)
		if err != nil {
			fmt.Println(function, err)
		}
		fmt.Println(function, "otp is invalid")
		return domain.Response{Code: "453", Msg: "otp is invalid"}
	}

	otpDetails.Validated = 1
	updateColumns = append(updateColumns, "validated")
	err = b.userPersistence.UserOtpPersistence.UpdateYpUserOtpByColumn(otpDetails, updateColumns...)
	if err != nil {
		fmt.Println(function, err)
	}
	response.Token = otpRequest.Token
	if userDetails.Mobile == "" {
		response.UserState = "kyc"
		response.Auth = userDetails.Auth
	} else {
		response.UserState = "home"
		response.CustomToken = "cuebdyucvdcbudbcsdbcdbcudbcb"
	}

	return domain.Response{Code: "200", Msg: "success", Model: response}
}

func (b *Login) AddUserOtpToDB(uid int, otp string, sentTo string, token string) error {
	newOtp := &domain.UserOtp{
		Uid:       uid,
		SentOn:    time.Now(),
		Validity:  time.Now().Add(time.Minute * 5),
		Otp:       otp,
		Token:     token,
		SentTo:    sentTo,
		Tries:     0,
		Validated: 0,
		UpdatedOn: time.Now(),
		CreatedOn: time.Now(),
	}
	_, err := b.userPersistence.UserOtpPersistence.AddYUserOtp(newOtp)
	return err
}
