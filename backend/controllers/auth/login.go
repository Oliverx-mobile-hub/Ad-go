package auth

import (
	"backend/config"
	"backend/utils/jwtutil"
	"backend/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录
func Login(r *gin.Context) {
	logs.Debug(nil, "登录")
	//一、从前端获取用户名和密码
	userInfo1 := UserInfo{}
	returnData := config.NewReturnDate()
	if err := r.ShouldBindJSON(&userInfo1); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		r.JSON(returnData.Status, returnData)
		return
	}
	logs.Debug(map[string]interface{}{"username": userInfo1.Username, "password": userInfo1.Password}, "准备验证登录信息")
	//二、验证用户名和密码
	if userInfo1.Username == config.Username && userInfo1.Password == config.Password {
		//2.1、认证成功，生成jwt token
		ss, err := jwtutil.GenToken(userInfo1.Username)
		if err != nil {
			logs.Error(map[string]interface{}{"用户名": userInfo1.Username, "错误信息": err.Error()},
				"用户名密码正确,但生成token失败")
			r.JSON(200, gin.H{
				"status":  401,
				"message": "生成token失败",
			})
			return
		}
		//token正常生成，返回给前端
		logs.Info(map[string]interface{}{"用户名": userInfo1.Username}, "登录成功")
		returnData.Status = 200
		returnData.Message = "登录成功"
		returnData.Data["token"] = ss
		returnData.Data = map[string]interface{}{
			"token": ss,
		}
		r.JSON(returnData.Status, returnData)
		return
	} else {
		//2.2、用户名和密码错误
		returnData.Status = 401
		returnData.Message = "用户名或密码错误"
		r.JSON(returnData.Status, returnData)
		return
	}
}

// 退出逻辑
func Logout(r *gin.Context) {
	// 退出
	// 实现退出逻辑
	r.JSON(200, gin.H{
		"message": "退出成功",
		"status":  200,
	})
	logs.Debug(nil, "用户已退出")
}
