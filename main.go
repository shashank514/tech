package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id    int    `orm:"column(id);auto"`
	Email string `orm:"column(email);null"`
	Fname string `orm:"column(fname);null"`
	Sname string `orm:"column(sname);null"`
}

func (t *User) TableName() string { return "yp_user" }
func init() {
	// Register the MySQL driver orm.RegisterDriver("mysql", orm.DRMySQL)
	// Register the default database
	orm.RegisterDataBase("default", "mysql", "root:jjhWmrXthSYGrWaoPelxRwzuzklCBwJG@tcp(monorail.proxy.rlwy.net:25594)/railway?charset=utf8")
	// Register model
	orm.RegisterModel(new(User))
	// Create table
	orm.RunSyncdb("default", false, false)
}
func main() {
	o := orm.NewOrm()
	// Insert data
	// user := User{Email: "Alice", Fname: "30"}
	// id, err := o.Insert(&user)
	// if err != nil {
	// fmt.Println("Error inserting data:", err)
	// } else {
	// fmt.Println("Data inserted with ID:", id) // }
	// Read data
	var users []User
	num, err := o.QueryTable("yp_user").All(&users)
	if err != nil {
		fmt.Println("Error reading data:", err)
	} else {
		fmt.Printf("Read %d users\n", num)
		for _, user := range users {
			fmt.Printf("ID: %d, Name: %s, Age: %s\n", user.Id, user.Email, user.Fname)
		}
	}
}
