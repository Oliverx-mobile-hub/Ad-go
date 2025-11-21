package display

import (
	"backend/utils/logs"

	"github.com/gin-gonic/gin"
)

// 由于每个接口都需要连接数据库，所以封装一个函数：GetDBWithErrorHandling，方便调用，并返回数据库连接对象
func GetDBWithErrorHandling(r *gin.Context) (*MysqlStore, bool) {
	db, err := NewMysqlStore()
	if err != nil {
		logs.Error(nil, "初始化mysql存储失败: "+err.Error())
		r.JSON(500, gin.H{"error": "数据库初始化失败", "details": err.Error()})
		return nil, false
	}
	return db, true
}
