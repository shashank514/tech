package login

import (
	"context"
	"fmt"
	"github.com/tech/core/domain"
	"time"
)

// GetUserPersonalDetails get user personal details
func (b *Login) GetUserPersonalDetails(exeCtx context.Context, user *domain.User, personalDetails *domain.PersonalDetails) domain.Response {
	funcName := "GetUserPersonalDetails"

	user.Name = personalDetails.Name
	user.Mobile = personalDetails.Mobile
	user.Gender = personalDetails.Gender
	user.Age = personalDetails.Age
	user.Profession = personalDetails.Profession
	user.Status = "home"

	updateColumns := []string{"status", "mobile", "name", "gender", "age", "profession"}
	err := b.userPersistence.YpUserPersistence.UpdateYpUserByColumn(user, updateColumns...)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: "error while updating user details"}
	}
	return domain.Response{Code: "200", Msg: "success"}
}

// GetUserAddressDetails get user Address
func (b *Login) GetUserAddressDetails(exeCtx context.Context, user *domain.User, userAddress *domain.AddressDetails) domain.Response {
	funcName := "GetUserAddressDetails"

	newAddress := &domain.UserAddress{
		Uid:       user.Id,
		Line:      userAddress.Line,
		City:      userAddress.City,
		State:     userAddress.State,
		Pincode:   userAddress.Pincode,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}

	_, err := b.userPersistence.UserAddressPersistence.AddYpUserAddress(newAddress)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: "error while updating user details"}
	}

	user.Status = "home"

	updateColumns := []string{"status"}
	err = b.userPersistence.YpUserPersistence.UpdateYpUserByColumn(user, updateColumns...)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: "error while updating user details"}
	}

	return domain.Response{Code: "200", Msg: "success"}
}

func (b *Login) GetUserDetails(exeCtx context.Context, user *domain.User) domain.Response {
	response := domain.UserDetails{}
	response.Status = user.Status
	return domain.Response{Code: "200", Msg: "success", Model: response}
}
