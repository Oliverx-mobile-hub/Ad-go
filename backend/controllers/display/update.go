package display

import (
	"backend/config"
	"backend/utils/logs"
	"net/http"
	"os"
	"path/filepath"

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

	//4.1、验证新的图片URL对应的文件是否存在
	if updateData.ImageURL != "" && updateData.ImageURL != existingImage.ImageURL {
		// 从URL中提取文件名,比如1.jpg
		filename := filepath.Base(updateData.ImageURL)
		//fmt.Println(filename)
		// 使用环境变量配置的uploads目录路径，比如/Users/zm/Desktop/Ad-go/backend/uploads/1.jpg
		filePath := filepath.Join(config.UploadsDir, filename)
		//fmt.Println(filePath)

		// 检查文件是否存在
		if _, err := os.Stat(filePath); err != nil {
			logs.Error(nil, "新的图片文件不存在: "+filePath)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "新的图片文件不存在于uploads目录中,请检查是否上传成功",
			})
			return
		}

		// 4.2、检查新的图片URL是否已被其他记录使用
		// 允许更换图片：如果选择了一个已被其他记录使用的图片，实际上是"交换图片"操作
		// 需要找到使用该图片的记录，并将其图片URL更新为当前记录的旧URL
		images, err := db.GetAllImages()
		if err != nil {
			logs.Error(nil, "获取所有图片失败")
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "系统内部错误",
			})
			return
		}

		// 查找是否已有其他记录使用了这个新的图片URL
		var existingRecordUsingNewURL *config.UploadRequest
		for _, img := range images {
			if img.ID != id && img.ImageURL == updateData.ImageURL {
				existingRecordUsingNewURL = &img
				break
			}
		}

		// 如果新的图片URL已被其他记录使用，执行图片交换操作
		if existingRecordUsingNewURL != nil {
			// 将其他记录的图片URL更新为当前记录的旧URL（实现图片交换）
			existingRecordUsingNewURL.ImageURL = existingImage.ImageURL
			if err := db.Update(existingRecordUsingNewURL); err != nil {
				logs.Error(nil, "更新其他记录失败: "+err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  500,
					"message": "图片交换操作失败",
				})
				return
			}
			logs.Info(nil, "成功交换图片: 记录"+existingRecordUsingNewURL.ID+"现在使用"+existingImage.ImageURL)
		}
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

	logs.Info(map[string]interface{}{
		"id":        existingImage.ID,
		"image_url": existingImage.ImageURL,
		"desc":      existingImage.Desc,
	}, "更新图片成功111")
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "图片更新成功",
		"data":    existingImage,
	})
}
