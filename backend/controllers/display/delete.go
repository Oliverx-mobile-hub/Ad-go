package display

import (
	"backend/config"
	"backend/utils/logs"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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
	image, err := db.GetImageByID(id)
	if err != nil {
		logs.Error(nil, "图片不存在")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "图片不存在",
			"error":   err.Error(),
		})
		return
	}

	//4、删除数据库记录
	if err := db.DeleteImageByID(id); err != nil {
		logs.Error(nil, "删除图片失败")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "删除图片失败",
			"error":   err.Error(),
		})
		return
	}

	//5、删除uploads文件夹中的实际文件
	if image.ImageURL != "" {
		// 从URL中提取文件名，如：1.jpg
		filename := filepath.Base(image.ImageURL)
		fmt.Println(filename)
		// 使用环境变量配置的uploads目录路径
		filePath := filepath.Join(config.UploadsDir, filename)

		// 检查文件是否存在并删除
		if _, err := os.Stat(filePath); err == nil {
			if err := os.Remove(filePath); err != nil {
				logs.Warning(nil, "删除文件失败: "+filePath)
			} else {
				logs.Info(nil, "成功删除文件: "+filePath)
			}
		} else {
			logs.Warning(nil, "文件不存在或无法访问: "+filePath)
		}
	}

	logs.Info(nil, "删除图片成功")
	c.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"message": "图片删除成功",
	})
}
