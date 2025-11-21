package auth

import (
	auth "backend/controllers/auth"

	"github.com/gin-gonic/gin"
)

// gin路由组的第二种写法,简洁，推荐第二种
func RegisterAuthRoutes(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.GET("/logout", auth.Logout)
	}
}
