package middleware

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func DBConnectionWithEnvMiddleware(c *gin.Context) {
	const functionName = "helper.DBConnectionWithEnvMiddleware"
	err := SetupMysqlConnections()
	if err != nil {
		fmt.Println(functionName, err)
		c.JSON(http.StatusInternalServerError, "DB Error")
		return
	}
	c.Next()
}

func SetupMysqlConnections() error {
	_, err := orm.GetDB("default")
	if err != nil {
		err = orm.RegisterDataBase("default", "mysql", "u374538722_sql:Sql@2023@tcp(srv687.hstgr.io:3306)/u374538722_newsql?charset=utf8", 30, 100)
	}
	SqlDebug := cast.ToInt(os.Getenv("SqlDebug"))
	if SqlDebug == 1 {
		orm.Debug = true
	}
	return err
}
