// 中间件层
package middlerwares

import (
	"backend/config"
	"backend/utils/jwtutil"
	"backend/utils/logs"

	"github.com/gin-gonic/gin"
)

// CORS 处理跨域请求
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有来源
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		// 允许携带凭证
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理OPTIONS请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func JWTAuth(r *gin.Context) {
	// 1. 除了login和logout之外的所有的接口，都要验证请求是否携带token，并且token是否合法
	requestUrl := r.Request.URL.Path
	logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "")
	// 不需要验证token的路径：登录、退出、静态文件
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" ||
		requestUrl == "/uploads" || requestUrl == "/uploads/" ||
		(len(requestUrl) > 8 && requestUrl[:8] == "/uploads") {
		logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "登录、退出和静态文件不需要验证token")
		r.Next()
		return
	}
	returnData := config.NewReturnDate()
	// token
	// 其他接口需要验证token
	// 获取是否携带token
	tokenString := r.Request.Header.Get("Authorization")
	if tokenString == "" {
		// 说明请求没有携带token
		returnData.Status = 401
		returnData.Message = "请求未携带Token, 请登录后尝试"
		r.JSON(returnData.Status, returnData)
		r.Abort()
		return
	}
	// token不为空，要去验证token是否合法
	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		returnData.Status = 401
		returnData.Message = "Token验证未通过"
		r.JSON(returnData.Status, returnData)
		r.Abort()
		return
	}
	// 验证成功
	r.Set("claims", claims)
	r.Next()
}
