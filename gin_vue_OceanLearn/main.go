package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 随机生成字符串
func RandomString(n int) string {
	var letters = []byte("salkdchsakjfbgcasbckjhwgahkfdiqwrfghwejfbsdjcvbczxczxcvv")
	result := make([]byte, n)

	//随机数种子
	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

// 数据库连接
func sql_conn() (engine *xorm.Engine) {
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

	err = engine.Sync(new(go_gin_User))

	if err != nil {
		fmt.Println("表结果同步失败")

	}
	return engine
}

//判断手机号是否存在

func isTelephoneExist(engine *xorm.Engine, telephone string) bool {
	var user go_gin_User
	engine.Where("telephone=?", telephone).Get(&user)

	if user.Id != 0 {
		return true
	}

	return false
}

type go_gin_User struct {
	Id        int
	Name      string `xorm:"varchar(20)"`
	Telephone string `xorm:"varchar(11)"`
	Password  string `xorm:"varchar(20)"`
}

func main() {

	//数据库连接
	engine := sql_conn()
	//当程序结束时关闭数据库连接
	defer engine.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {

		//获取前端返回参数
		name := ctx.PostForm("name")

		telephone := ctx.PostForm("telephone")

		password := ctx.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须是11位",
			})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码需要大于等于6位",
			})
			return
		}

		//若未传递名称，则给一个10位的随机字符串

		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, password, telephone)
		//判断手机号是否存在
		if isTelephoneExist(engine, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户已存在",
			})
			return

		}
		//创建用户
		newUser := go_gin_User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		// 向数据库插入数据
		engine.Insert(&newUser)

		//返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})

	})
	r.Run(":8234")

}
