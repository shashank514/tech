package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YpInvestmentBuyDetails struct {
	Id             int       `orm:"column(id);auto"`
	Uid            int       `orm:"column(uid);null"`
	Date           int       `orm:"column(date);null"`
	Month          string    `orm:"column(month);null"`
	Year           int       `orm:"column(year);null"`
	Type           string    `orm:"column(type);null"`
	Name           string    `orm:"column(name);null"`
	Symbol         string    `orm:"column(symbol);null"`
	Enable         int       `orm:"column(enable);null"`
	BuyCount       string    `orm:"column(buy_count);null"`
	AmountPerBuy   string    `orm:"column(amount_per_buy);null"`
	TotalAmount    string    `orm:"column(total_amount);null"`
	RemainingCount string    `orm:"column(remaining_count);null"`
	FdInterest     string    `orm:"column(fd_interest);null"`
	UpdatedOn      time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn      time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}

func (t *YpInvestmentBuyDetails) TableName() string {
	return "yp_investment_buy_details"
}

func init() {
	orm.RegisterModel(new(YpInvestmentBuyDetails))
}

func (t *YpInvestmentBuyDetails) AddYpInvestmentBuyDetails() (int64, error) {
	return orm.NewOrm().Insert(t)
}

func (t *YpInvestmentBuyDetails) GetAllYpInvestmentBuyDetailsByUid(uid int) (data []*YpInvestmentBuyDetails, err error) {
	o := orm.NewOrm()
	data = []*YpInvestmentBuyDetails{}
	_, err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("enable", 1).All(&data)
	return
}

func (t *YpInvestmentBuyDetails) GetInvestmentBuyDetailsByType(types string, uid int) (data []*YpInvestmentBuyDetails, err error) {
	o := orm.NewOrm()
	data = []*YpInvestmentBuyDetails{}
	_, err = o.QueryTable(t.TableName()).Filter("uid", uid).Filter("type", types).Filter("enable", 1).All(&data)
	return
}
