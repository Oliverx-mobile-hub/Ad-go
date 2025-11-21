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

	r.Use(middlerwares.CORS())  //启用CORS中间件
	r.Use(middlerwares.JWTAuth) //启用中间件，验证token，其他接口验证token后，才可以使用
	logs.Info(nil, "程序启动")

	// 添加静态文件服务，让前端可以访问上传的图片
	r.Static("/uploads", "./uploads")

	//注册路由
	router.RegisterRoutes(r) //总入口
	//启动服务
	r.Run(":" + config.Port) //端口号
}
