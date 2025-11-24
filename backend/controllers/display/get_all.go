package display

import (
	"net/http"

	"backend/utils/logs"

	"github.com/gin-gonic/gin"
)

// GetAllImages 获取所有图片
func GetAllImages(c *gin.Context) {
	db, err := NewMysqlStore()
	if err != nil {
		logs.Error(nil, "初始化数据库失败")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "连接数据库失败",
			"error":   err.Error(),
		})
		return
	}

	images, err := db.GetAllImages()
	if err != nil {
		logs.Error(nil, "获取图片列表失败")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "获取图片列表失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "获取图片列表成功",
		"data":    images,
	})
}
