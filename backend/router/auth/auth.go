package auth

import (
	auth "backend/controllers/auth"

	"github.com/gin-gonic/gin"
)

// 公共认证路由 - 不需要认证
func RegisterPublicAuthRoutes(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	{
		// 公共接口：登录和登出不需要认证
		authGroup.POST("/login", auth.Login)
		authGroup.GET("/logout", auth.Logout)
	}
}
