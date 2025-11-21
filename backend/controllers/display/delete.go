package display

import (
	"backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteImage(c *gin.Context) {
	logs.Info(nil, "删除图片信息")

	//1、连接数据库
	db, ok := GetDBWithErrorHandling(c)
	if !ok {
		return
	}

	//2、获取图片ID
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "图片ID不能为空",
		})
		return
	}

	//3、检查图片是否存在
	_, err := db.GetImageByID(id)
	if err != nil {
		logs.Error(nil, "图片不存在")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "图片不存在",
			"error":   err.Error(),
		})
		return
	}

	//4、删除图片
	if err := db.DeleteImageByID(id); err != nil {
		logs.Error(nil, "删除图片失败")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "删除图片失败",
			"error":   err.Error(),
		})
		return
	}

	logs.Info(nil, "删除图片成功")
	c.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"message": "图片删除成功",
	})
}
