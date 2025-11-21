package display

import (
	"backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetImageByID 根据ID获取单张图片
func GetImageByID(c *gin.Context) {
	id := c.Param("id") //推荐用param，路径明确，参数也明确
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "图片ID不能为空",
		})
		return
	}

	//连接数据库
	db, ok := GetDBWithErrorHandling(c)
	if !ok {
		return
	}

	images1, err := db.GetImageByID(id)
	if err != nil {
		logs.Error(map[string]interface{}{"id": id}, "获取图片失败")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "图片不存在",
			"error":   err.Error(),
		})
		return
	}

	logs.Info(map[string]interface{}{"id": id}, "获取图片成功")
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "获取图片成功",
		"data":    images1,
	})
}

// GetAllImages 获取所有图片
// func GetAllImages(c *gin.Context) {
// 	db, err := NewMysqlStore()
// 	if err != nil {
// 		logs.Error(nil, "初始化数据库失败")
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  500,
// 			"message": "连接数据库失败",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	images, err := db.GetAllImages()
// 	if err != nil {
// 		logs.Error(nil, "获取图片列表失败")
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  500,
// 			"message": "获取图片列表失败",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"message": "获取图片列表成功",
// 		"data":    images,
// 	})
// }
