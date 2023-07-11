package middleware

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header

		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})

			//抛弃请求
			ctx.Abort()

			return
		}
		//截取token有效位
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足2",
			})
			//抛弃请求
			ctx.Abort()

			return
		}

		//验证通过后获取claim中的userId
		userId := claims.UserId

		engine := common.GetEngine()

		var user model.Go_gin_User

		engine.Where("id=?", userId).Get(&user)

		if user.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足3",
			})
			//抛弃请求
			ctx.Abort()

			return
		}

		//用户存在 将user信息写入上下文

		ctx.Set("user", user)

		ctx.Next()

	}
}
