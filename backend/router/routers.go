package router

import (
	auth "backend/router/auth"
	display "backend/router/display"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterAuthRoutes(apiGroup)       //auth子路由配置
	display.RegisterDisplayRoutes(apiGroup) //display子路由配置
}
