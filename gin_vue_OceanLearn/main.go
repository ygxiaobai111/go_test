package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"OceanLearn/common"
	_ "OceanLearn/model"
)

func main() {

	//数据库连接
	engine := common.Sql_conn()
	//当程序结束时关闭数据库连接
	defer engine.Close()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":8234")

}
