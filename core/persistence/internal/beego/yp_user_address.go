package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoUserAddress struct{}

func (t *BeegoUserAddress) AddYpUserAddress(address *domain.UserAddress) (id int64, err error) {
	v := t.convertDomainToModel(address)
	return v.AddYpUserAddress()
}

func (t *BeegoUserAddress) convertDomainToModel(data *domain.UserAddress) *ypmodel.YPUserAddress {
	return &ypmodel.YPUserAddress{
		Id:        data.Id,
		Uid:       data.Uid,
		Line:      data.Line,
		City:      data.City,
		District:  data.District,
		Pincode:   data.Pincode,
		UpdatedOn: data.UpdatedOn,
		CreatedOn: data.CreatedOn,
	}
}

func (t *BeegoUserAddress) convertModelToDomain(data *ypmodel.YPUserAddress) *domain.UserAddress {
	return &domain.UserAddress{
		Id:        data.Id,
		Uid:       data.Uid,
		Line:      data.Line,
		City:      data.City,
		District:  data.District,
		Pincode:   data.Pincode,
		UpdatedOn: data.UpdatedOn,
		CreatedOn: data.CreatedOn,
	}
}
