package beego

import (
	"github.com/tech/core/domain"
	"github.com/tech/model/ypmodel"
)

type BeegoYpUser struct{}

func (t *BeegoYpUser) AddYPUser(user *domain.User) (id int64, err error) {
	v := t.convertDomainToModel(user)
	return v.AddYPUser()
}

func (t *BeegoYpUser) GetYPUserByEmail(email string) (user *domain.User, err error) {
	data, err := new(ypmodel.YPUser).GetYPUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return t.convertModelToDomain(data), err
}

func (t *BeegoYpUser) convertDomainToModel(data *domain.User) *ypmodel.YPUser {
	return &ypmodel.YPUser{
		Id:        data.Id,
		Auth:      data.Auth,
		Email:     data.Email,
		Mobile:    data.Mobile,
		Name:      data.Name,
		CreatedOn: data.CreatedOn,
		UpdatedOn: data.UpdatedOn,
	}
}

func (t *BeegoYpUser) convertModelToDomain(data *ypmodel.YPUser) *domain.User {
	return &domain.User{
		Id:        data.Id,
		Auth:      data.Auth,
		Email:     data.Email,
		Mobile:    data.Mobile,
		Name:      data.Name,
		CreatedOn: data.CreatedOn,
		UpdatedOn: data.UpdatedOn,
	}
}
