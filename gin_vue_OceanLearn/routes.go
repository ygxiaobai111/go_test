package main

import (
	"OceanLearn/controller"
	"OceanLearn/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	//AuthMiddleware()为验证中间件
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
