package display

import (
	upload_image "backend/config"
	"backend/utils/logs"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"backend/config"
)

func Upload(c *gin.Context) {
	logs.Info(nil, "开始上传图片")

	//1、连接数据库
	db, ok := GetDBWithErrorHandling(c)
	if !ok {
		return
	}

	//2.1、获取图片ID数据
	id := c.PostForm("id")
	if id == "" {
		logs.Error(nil, "图片ID不能为空")
		c.JSON(400, gin.H{
			"status":  400,
			"message": "图片ID不能为空",
		})
		return
	}

	//2.2、处理文件上传 (multipart/form-data格式)
	file, err := c.FormFile("images")
	if err != nil {
		logs.Error(nil, "获取上传文件失败")
		c.JSON(400, gin.H{
			"status":  400,
			"message": "请选择要上传的图片文件",
			"error":   err.Error(),
		})
		return
	}

	//2.3、获取其他表单参数
	desc := c.PostForm("desc")
	if desc == "" {
		logs.Error(nil, "描述信息不能为空")
		c.JSON(400, gin.H{
			"status":  400,
			"message": "图片描述信息不能为空",
		})
		return
	}

	//4、保存上传的文件，数据库是索引，/Users/zm/Desktop/Ad-go/backend/uploads是真实目录，两者缺一不可。
	uploadDir := config.UploadsDir
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		logs.Error(nil, "创建上传目录失败")
		c.JSON(500, gin.H{
			"status":  500,
			"message": "服务器内部错误",
			"error":   err.Error(),
		})
		return
	}

	fileExt := filepath.Ext(file.Filename)         //filepath.Ext() 函数作用：获取上传文件的扩展名，如dog.jpg，fileExt为.jpg
	fileName := fmt.Sprintf("%s%s", id, fileExt)   //id=1,fileExt=.jpg,fileName=1.jpg
	filePath := filepath.Join(uploadDir, fileName) //uploadDir=./uploads,fileName=1.jpg,filePath=./uploads/1.jpg

	// 4.1、保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logs.Error(nil, "保存文件失败")
		c.JSON(500, gin.H{
			"status":  500,
			"message": "保存图片文件失败",
			"error":   err.Error(),
		})
		return
	}

	//5、生成可访问的URL
	imageURL1 := fmt.Sprintf("/uploads/%s", fileName)

	//5.1、构建请求对象
	UploadImage := upload_image.UploadRequest{
		ID:       id,
		ImageURL: imageURL1,
		Desc:     desc,
	}

	//5.2、上传图片到数据库
	if err := db.Upload(&UploadImage); err != nil {
		logs.Error(nil, "上传图片到数据库失败")
		c.JSON(500, gin.H{
			"status":  500,
			"message": "上传图片到数据库失败",
			"error":   err.Error(),
		})
		return
	}

	//5.3、打印上传图片信息日志
	logs.Info(map[string]interface{}{
		"id":        UploadImage.ID,
		"image_url": UploadImage.ImageURL,
		"desc":      UploadImage.Desc,
	}, "上传图片成功")

	//6、返回上传图片信息
	c.JSON(200, gin.H{
		"status":  200,
		"message": "图片上传成功",
		"data":    UploadImage,
	})
}
