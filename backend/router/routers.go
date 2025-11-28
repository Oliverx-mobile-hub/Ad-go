package router

import (
	"backend/middlerwares"
	auth "backend/router/auth"
	display "backend/router/display"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 公共API路由组 - 不需要认证
	publicApiGroup := r.Group("/api")
	{
		// 注册公共路由（不需要认证）
		auth.RegisterPublicAuthRoutes(publicApiGroup)
		display.RegisterPublicDisplayRoutes(publicApiGroup)
	}

	// 受保护API路由组 - 需要认证（应用JWT中间件）
	protectedApiGroup := r.Group("/api")
	protectedApiGroup.Use(middlerwares.JWTAuth) // 为受保护路由组应用JWT中间件
	{
		// 注册受保护路由（需要认证）
		display.RegisterProtectedDisplayRoutes(protectedApiGroup)
	}
}
