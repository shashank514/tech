package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YpStockName struct {
	Id        int       `orm:"column(id);auto"`
	Enable    int       `orm:"column(enable);null"`
	StockName string    `orm:"column(stock_name);null"`
	Symbol    string    `orm:"column(symbol);null"`
	Price     float64   `orm:"column(price);null"`
	Category  string    `orm:"column(category);null"`
	Sector    string    `orm:"column(sector);null"`
	Industry  string    `orm:"column(industry);null"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}

func (t *YpStockName) TableName() string {
	return "yp_stock_name"
}

func init() {
	orm.RegisterModel(new(YpStockName))
}

func (t *YpStockName) AddYpStockName() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)
	return
}

func (t *YpStockName) GetAllYpStockName() (list []*YpStockName, err error) {
	o := orm.NewOrm()
	list = []*YpStockName{}
	_, err = o.QueryTable(t.TableName()).All(&list)
	return
}

func (t *YpStockName) GetAllYpDetailsNameByCategory(category string) (list []*YpStockName, err error) {
	o := orm.NewOrm()
	list = []*YpStockName{}
	_, err = o.QueryTable(t.TableName()).Filter("enable", 1).Filter("category", category).All(&list)
	return
}

func (t *YpStockName) GetYpStockNameByName(name string) (v *YpStockName, err error) {
	o := orm.NewOrm()
	v = &YpStockName{}
	err = o.QueryTable(t.TableName()).Filter("stock_name", name).One(v)
	return
}

func (t *YpStockName) GetYpStockNameBySymbol(symbol string) (v *YpStockName, err error) {
	o := orm.NewOrm()
	v = &YpStockName{}
	err = o.QueryTable(t.TableName()).Filter("symbol", symbol).One(v)
	return
}
