package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoUserOtp struct{}

func (t *BeegoUserOtp) AddYUserOtp(otp *domain.UserOtp) (id int64, err error) {
	v := t.convertDomainToModel(otp)
	return v.AddYUserOtp()
}

func (t *BeegoUserOtp) GetYpUserOtpByToken(token string) (otp *domain.UserOtp, err error) {
	v, err := new(ypmodel.YpUserOtp).GetYpUserOtpByToken(token)
	if err != nil {
		return nil, err
	}
	return t.convertModelToDomain(v), nil
}

func (t *BeegoUserOtp) UpdateYpUserOtpByColumn(data *domain.UserOtp, column ...string) (err error) {
	v := t.convertDomainToModel(data)
	err = v.UpdateYpUserOtpByColumn(column...)
	return err
}

func (t *BeegoUserOtp) convertDomainToModel(data *domain.UserOtp) *ypmodel.YpUserOtp {
	return &ypmodel.YpUserOtp{
		Id:        data.Id,
		Uid:       data.Uid,
		SentTo:    data.SentTo,
		Validity:  data.Validity,
		Otp:       data.Otp,
		Token:     data.Token,
		SentOn:    data.SentOn,
		Tries:     data.Tries,
		Validated: data.Validated,
		CreatedOn: data.CreatedOn,
		UpdatedOn: data.UpdatedOn,
	}
}

func (t *BeegoUserOtp) convertModelToDomain(data *ypmodel.YpUserOtp) *domain.UserOtp {
	return &domain.UserOtp{
		Id:        data.Id,
		Uid:       data.Uid,
		SentTo:    data.SentTo,
		Validity:  data.Validity,
		Otp:       data.Otp,
		Token:     data.Token,
		SentOn:    data.SentOn,
		Tries:     data.Tries,
		Validated: data.Validated,
		CreatedOn: data.CreatedOn,
		UpdatedOn: data.UpdatedOn,
	}
}
