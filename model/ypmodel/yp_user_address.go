package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YPUserAddress struct {
	Id        int       `orm:"column(id);auto"`
	Uid       int       `orm:"column(uid);null"`
	Line      string    `orm:"column(line);null"`
	City      string    `orm:"column(city);null"`
	District  string    `orm:"column(district);null"`
	State     string    `orm:"column(state);null"`
	Pincode   string    `orm:"column(pincode);null"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}

func (t *YPUserAddress) TableName() string {
	return "yp_user_address"
}

func init() {
	orm.RegisterModel(new(YPUserAddress))
}

func (t *YPUserAddress) AddYpUserAddress() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(t)
	return id, err
}
