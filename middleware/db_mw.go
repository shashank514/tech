package middleware

import (
	"fmt"
	"net/http"

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
	err := orm.RegisterDataBase("default", "mysql", "u374538722_sql:Sql@2023@tcp(srv687.hstgr.io:3306)/u374538722_newsql?charset=utf8", 30, 100)
	return err
}
