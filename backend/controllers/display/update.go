package display

import (
	"backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	logs.Info(nil, "更新图片信息")

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
	existingImage, err := db.GetImageByID(id)
	if err != nil {
		logs.Error(nil, "图片不存在")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "图片不存在",
			"error":   err.Error(),
		})
		return
	}

	//4、绑定更新数据
	var updateData struct {
		ImageURL string `json:"image_url"`
		Desc     string `json:"desc"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		logs.Error(nil, "绑定JSON失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "请求参数格式错误",
			"error":   err.Error(),
		})
		return
	}

	//5、更新图片信息
	existingImage.ImageURL = updateData.ImageURL
	existingImage.Desc = updateData.Desc

	if err := db.Update(existingImage); err != nil {
		logs.Error(nil, "更新图片失败: %v")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "更新图片失败",
			"error":   err.Error(),
		})
		return
	}

	logs.Info(nil, "更新图片成功,ID: %s")
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "图片更新成功",
		"data":    existingImage,
	})
}
