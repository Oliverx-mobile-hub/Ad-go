package main

import (
	"backend/router"
	"backend/utils/logs"

	"backend/config"

	"backend/middlerwares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //gin引擎

	r.Use(middlerwares.CORS()) //启用CORS中间件（所有路由都需要）
	logs.Info(nil, "程序启动")

	// 添加静态文件服务，让前端可以访问上传的图片
	r.Static("/uploads", "./uploads")

	// 注册路由（分公共路由&受保护路由）
	router.RegisterRoutes(r) //总入口

	//启动服务
	r.Run(":" + config.Port) //端口号
}
