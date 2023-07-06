package common

import (
	"OceanLearn/model"
	"fmt"
	"log"

	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
)

// 数据库连接
func Sql_conn() (engine *xorm.Engine) {
	//数据库连接基本信息
	var (
		userName  string = "root"
		password  string = ""
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)
	//构建数据库连接信息
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)

	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatal("数据库连接失败")
	}

	err = engine.Sync(new(model.Go_gin_User))

	if err != nil {
		fmt.Println("表结果同步失败")

	}
	Engine = engine
	return engine
}

func GetEngine() *xorm.Engine {
	return Engine
}
