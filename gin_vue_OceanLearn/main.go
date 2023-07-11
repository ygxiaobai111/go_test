package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"OceanLearn/common"
	_ "OceanLearn/model"
)

func main() {

	//引入配置文件
	InitConfig()
	//数据库连接
	engine := common.Sql_conn()
	//当程序结束时关闭数据库连接
	defer engine.Close()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}

}

// 连接.yml配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	//读取文件名
	viper.SetConfigName("application")

	viper.SetConfigType("yml")

	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
