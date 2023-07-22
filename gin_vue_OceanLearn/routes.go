package main

import (
	"OceanLearn/controller"
	"OceanLearn/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/index", controller.Index)
	r.GET("/register", controller.RegPage)
	r.GET("/login", controller.LogPage)
	r.GET("/api/auth/dict", controller.DictPage)
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	//AuthMiddleware()为验证中间件
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	//词典
	r.POST("/api/auth/dict", controller.Dict)

	//404
	r.NoRoute(func(ctx *gin.Context) { ctx.HTML(404, "404.html", nil) })
	return r
}
