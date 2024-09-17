package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YpExpenseDate struct {
	Id        int       `orm:"column(id);auto"`
	Uid       int       `orm:"column(uid);null"`
	Date      int       `orm:"column(date);null"`
	Month     string    `orm:"column(month);null"`
	Year      int       `orm:"column(year);null"`
	Amount    string    `orm:"column(amount);null"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}

func (t *YpExpenseDate) TableName() string {
	return "yp_expense_date"
}

func init() {
	orm.RegisterModel(new(YpExpenseDate))
}

func (t *YpExpenseDate) AddYpExpenseDate() (id int64, err error) {
	o := orm.NewOrm()
	return o.Insert(t)
}

func (t *YpExpenseDate) GetYpExpenseDateById(uid int, month string, year int) (v []*YpExpenseDate, err error) {
	o := orm.NewOrm()
	v = []*YpExpenseDate{}
	_, err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("month", month).Filter("year", year).All(&v)
	return v, err
}
