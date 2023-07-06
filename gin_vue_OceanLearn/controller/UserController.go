package controller

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"OceanLearn/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

//判断手机号是否存在

func isTelephoneExist(engine *xorm.Engine, telephone string) bool {
	var user model.Go_gin_User
	engine.Where("telephone=?", telephone).Get(&user)

	if user.Id != 0 {
		return true
	}

	return false
}

func Register(ctx *gin.Context) {
	engine := common.GetEngine()
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
		name = util.RandomString(10)
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
	newUser := model.Go_gin_User{
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

}
