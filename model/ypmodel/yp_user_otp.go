package ypmodel

import "time"

type YUserOtp struct {
	Id        int       `orm:"column(id);auto"`
	Uid       string    `orm:"column(uid);null"`
	SentOn    time.Time `orm:"column(sentOn);type(datetime);null"`
	Validity  time.Time `orm:"column(validity);type(datetime);null"`
	Otp       string    `orm:"column(otp);null"`
	Token     string    `orm:"column(token);size(32);null"`
	SentTo    string    `orm:"column(sentTo);size(64);null"`
	Tries     int       `orm:"column(tries);null"`
	Validated int8      `orm:"column(validated);null"`
	UpdatedOn time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
	CreatedOn time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
}
