package ypmodel

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type YpUserOtp struct {
	Id        int       `orm:"column(id);auto"`
	Uid       int       `orm:"column(uid);null"`
	SentOn    time.Time `orm:"column(sentOn);type(datetime);null"`
	Validity  time.Time `orm:"column(validity);type(datetime);null"`
	Otp       string    `orm:"column(otp);null"`
	Token     string    `orm:"column(token);size(32);null"`
	SentTo    string    `orm:"column(sentTo);size(64);null"`
	Validated int       `orm:"column(validated);null"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
	Tries     int       `orm:"column(tries);null"`
}

func (t *YpUserOtp) TableName() string {
	return "yp_user_otp"
}

func init() {
	orm.RegisterModel(new(YpUserOtp))
}

func (t *YpUserOtp) AddYUserOtp() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)
	return
}

func (t *YpUserOtp) GetYpUserOtpByToken(token string) (otp *YpUserOtp, err error) {
	o := orm.NewOrm()
	otp = &YpUserOtp{}
	err = o.QueryTable(t.TableName()).Filter("token", token).One(otp)
	return
}

func (t *YpUserOtp) UpdateYpUserOtpByColumn(columns ...string) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(t, columns...)
	return
}
