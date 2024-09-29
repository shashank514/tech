package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YPInvestmentDetails struct {
	Id                      int       `orm:"column(id);auto"`
	Uid                     int       `orm:"column(uid);null"`
	TotalInvestmentAmount   string    `orm:"column(total_investment_amount);null"`
	PresentInvestmentAmount string    `orm:"column(present_investment_amount);null"`
	ProfitAfterSellAmount   string    `orm:"column(profit_after_sell_amount);null"`
	LossAfterSellAmount     string    `orm:"column(loss_after_sell_amount);null"`
	UpdatedOn               time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn               time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}

func (t *YPInvestmentDetails) TableName() string {
	return "yp_investment_details"
}

func init() {
	orm.RegisterModel(new(YPInvestmentDetails))
}

func (t *YPInvestmentDetails) AddYpInvestmentDetails() (int64, error) {
	return orm.NewOrm().Insert(t)
}

func (t *YPInvestmentDetails) GetUserInvestmentDetailsByUid(uid int) (v *YPInvestmentDetails, err error) {
	o := orm.NewOrm()
	v = &YPInvestmentDetails{}
	err = o.QueryTable(t.TableName()).Filter("uid", uid).One(v)
	return
}

func (t *YPInvestmentDetails) UpdateYpInvestmentDetailsByColumns(columns ...string) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(t, columns...)
	return
}
