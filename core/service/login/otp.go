package login

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
	"github.com/tech/cloud/message"
	"github.com/tech/core/domain"
	"github.com/tech/core/persistence/user"
	"github.com/tech/core/service/login/driver"
	"github.com/tech/util"
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

	//check email domain
	if !isDomainValid(otpRequest.Email) {
		fmt.Println("invalid email")
		return domain.Response{Code: "420", Msg: "invalid email"}
	}

	// get user details from yp_user
	userDetails, err := b.userPersistence.YpUserPersistence.GetYPUserByEmail(otpRequest.Email)
	if err != nil && err != orm.ErrNoRows {
		fmt.Println(function, err)
		return domain.Response{Code: "452", Msg: "err while user details"}
	} else if err == orm.ErrNoRows {
		userDetails = &domain.User{
			Auth:   util.GenereateAuthToken(otpRequest.Email),
			Email:  otpRequest.Email,
			Status: "personalDetails",
		}

		id, err := b.userPersistence.YpUserPersistence.AddYPUser(userDetails)
		if err != nil {
			fmt.Println(function, err)
			return domain.Response{Code: "452", Msg: "err while creating user"}
		}
		userDetails.Id = cast.ToInt(id)
	}

	blockForTheDay, otp := b.CheckResendConditions(otpRequest.Email)
	if blockForTheDay {
		fmt.Println("user is block for the day to generate otp")
		return domain.Response{Code: "459", Msg: "otp is block"}
	}

	response.OtpDigits = otpDigits
	if otp == "" {
		otp = GenerateRandomOtp(otpDigits)
	}
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

	if otpDetails.Tries >= 3 {
		fmt.Println(function, "Submit invalid otp 3 times")
		return domain.Response{Code: "452", Msg: "Submit invalid otp 3 times generate otp"}
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
	response.UserState = userDetails.Status
	response.CustomToken = util.GenereateToken(userDetails.Auth)

	return domain.Response{Code: "200", Msg: "success", Model: response}
}

func (b *Login) CheckResendConditions(sentTo string) (bool, string) {
	var count int
	var row *domain.UserOtp
	today, _ := time.Parse("2006-01-02 15:04:05", todaysDate())
	count, row = b.userPersistence.UserOtpPersistence.GetYpUserOtpCount(sentTo, today)
	var blockForTheDay bool
	var oldOTP string

	// if there is no row, then return  false and empty for otp
	if count == 0 {
		return blockForTheDay, oldOTP
	}

	//if the OTP is already validated then generate a new OTP
	if row.Validated == 1 {
		return blockForTheDay, oldOTP
	}

	if count > cast.ToInt(3) {
		blockForTheDay = true
	}

	// if there are previous OTPs' and total OTP count is below the limit send the same OTP
	if count > 0 && count <= cast.ToInt(3) {
		oldOTP = cast.ToString(row.Otp)
		return blockForTheDay, oldOTP
	}
	return blockForTheDay, oldOTP
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
