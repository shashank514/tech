package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YPUser struct {
	Id        int       `orm:"column(id);auto"`
	Auth      string    `orm:"column(auth);null"`
	Email     string    `orm:"column(email);null"`
	Mobile    string    `orm:"column(mobile);null"`
	Name      string    `orm:"column(name);null"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);null;auto_now_add"`
}

func (t *YPUser) TableName() string {
	return "yp_user"
}

func init() {
	orm.RegisterModel(new(YPUser))
}

func (t *YPUser) AddYPUser() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)
	return id, err
}

func (t *YPUser) GetYPUserByEmail(email string) (data *YPUser, err error) {
	o := orm.NewOrm()
	data = &YPUser{}
	err = o.QueryTable(t.TableName()).Filter("email", email).One(data)
	return data, err
}
