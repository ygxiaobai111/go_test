package controller

import (
	"OceanLearn/common"
	"OceanLearn/dto"
	"OceanLearn/model"
	"OceanLearn/response"
	"OceanLearn/util"
	"context"
	"fmt"
	"log"
	"net/http"

	pb "OceanLearn/proto"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

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

// 首页
func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"msg":  "后台消息",
		"name": "标题是这样的",
	})
}

// 注册页面
func RegPage(ctx *gin.Context) {
	ctx.HTML(200, "register.html", gin.H{
		"msg": "后台消息",
	})
}

// 登录页面
func LogPage(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{
		"msg": "后台消息",
	})
}

func DictPage(ctx *gin.Context) {
	ctx.HTML(200, "dict.html", gin.H{
		"msg": "后台消息",
	})
}

// 注册
func Register(ctx *gin.Context) {
	engine := common.GetEngine()
	//获取前端返回参数
	name := ctx.PostForm("name")

	telephone := ctx.PostForm("telephone")

	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须是11位")

		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码需要大于等于6位")

		return
	}

	//若未传递名称，则给一个10位的随机字符串

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, password, telephone)
	//判断手机号是否存在
	if isTelephoneExist(engine, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")

		return

	}
	//创建用户

	//加密用户密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")

		return
	}
	newUser := model.Go_gin_User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	// 向数据库插入数据
	_, err = engine.Insert(&newUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "数据插入失败")
		return
	}
	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})

}

// 登录
func Login(ctx *gin.Context) {
	engine := common.GetEngine()
	//获取前端返回参数

	telephone := ctx.PostForm("telephone")

	password := ctx.PostForm("password")
	fmt.Println(telephone)
	//数据验证

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须是11位")

		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码需要大于等于6位")

		return
	}
	//判断手机号是否存在
	var user model.Go_gin_User
	engine.Where("telephone=?", telephone).Get(&user)
	if user.Id == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")

		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")

		return
	}
	//发放token

	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 500, nil, "系统异常")

		log.Printf("token generate error :%v", err)

		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登陆成功")

}

func Dict(ctx *gin.Context) {
	//数据库连接
	// engine := common.GetEngine()
	//获取前端返回参数

	inputText := ctx.PostForm("inputText")
	fmt.Println(inputText)
	conn := common.Grpc_conn()

	defer conn.Close()

	client := pb.NewSayDictClient(conn)

	resp, _ := client.SayDict(context.Background(), &pb.DictRequest{TransType: "en2zh", Source: inputText})
	data := dto.ToDictDto(resp)
	fmt.Println(resp.Dictionary)
	response.Success(ctx, gin.H{"data": data}, "返回成功")
	fmt.Println(data.Dictionary)

}

func Info(ctx *gin.Context) {

	fmt.Println("ok")
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{

			//将敏感信息隐藏
			"user": dto.ToUserDto(user.(model.Go_gin_User)),
		},
	})
}
