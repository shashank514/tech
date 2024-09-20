package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YpExpenseDetails struct {
	Id          int       `orm:"column(id);auto"`
	Uid         int       `orm:"column(uid);null"`
	Date        int       `orm:"column(date);null"`
	Month       string    `orm:"column(month);null"`
	Year        int       `orm:"column(year);null"`
	Amount      string    `orm:"column(amount);null"`
	Category    string    `orm:"column(category);null"`
	PaymentMode string    `orm:"column(payment_mode);null"`
	Description string    `orm:"column(description);null"`
	UpdatedOn   time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn   time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}

func (t *YpExpenseDetails) TableName() string {
	return "yp_expense_details"
}

func init() {
	orm.RegisterModel(new(YpExpenseDetails))
	orm.Debug = true
}

func (t *YpExpenseDetails) AddUserExpense() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)
	return
}

func (t *YpExpenseDetails) GetUserExpenseById(uid int, month string, year int) (v []*YpExpenseDetails, err error) {
	o := orm.NewOrm()
	v = []*YpExpenseDetails{}
	_, err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("month", month).Filter("year", year).All(&v)
	return v, err
}

func (t *YpExpenseDetails) GetUserExpenseByUidAndCategory(uid int, category, month string, year int) (v []*YpExpenseDetails, err error) {
	o := orm.NewOrm()
	v = []*YpExpenseDetails{}
	_, err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("category", category).Filter("month", month).Filter("year", year).All(&v)
	return v, err
}

func (t *YpExpenseDetails) GetUserExpenseByUidAndPaymentMode(uid int, paymentMode, month string, year int) (v []*YpExpenseDetails, err error) {
	o := orm.NewOrm()
	v = []*YpExpenseDetails{}
	_, err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("payment_mode", paymentMode).Filter("month", month).Filter("year", year).All(&v)
	return v, err
}
