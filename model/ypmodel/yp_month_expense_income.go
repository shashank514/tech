package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YpMonthIncomeExpense struct {
	Id             int       `orm:"column(id);auto"`
	Uid            int       `orm:"column(uid);null"`
	Month          string    `orm:"column(month);null"`
	Year           int       `orm:"column(year);null"`
	IncomeAmount   string    `orm:"column(income_amount);null"`
	ExpensesAmount string    `orm:"column(expenses_amount);null"`
	UpdatedOn      time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn      time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}

func (t *YpMonthIncomeExpense) TableName() string {
	return "yp_month_expense_income"
}

func init() {
	orm.RegisterModel(new(YpMonthIncomeExpense))
}

func (t *YpMonthIncomeExpense) AddYpMonthIncomeExpense() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)
	return
}

func (t *YpMonthIncomeExpense) GetYpMonthIncomeExpenseByYear(uid int, year int) (v []*YpMonthIncomeExpense, err error) {
	o := orm.NewOrm()
	v = []*YpMonthIncomeExpense{}
	_, err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("year", year).All(&v)
	return
}

func (t *YpMonthIncomeExpense) GetDetailsUsingUidAndMonth(uid int, month string, year int) (v *YpMonthIncomeExpense, err error) {
	o := orm.NewOrm()
	v = &YpMonthIncomeExpense{}
	err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("month", month).Filter("year", year).One(v)
	return
}

func (t *YpMonthIncomeExpense) UpdateYpMonthIncomeExpenseByColumn(columns ...string) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(t, columns...)
	return
}
