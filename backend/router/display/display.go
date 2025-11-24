package edit

import (
	"backend/controllers/display"

	"github.com/gin-gonic/gin"
)

// gin路由组：上传、更新图片路径
func RegisterDisplayRoutes(g *gin.RouterGroup) {
	displayGroup := g.Group("/display")
	{
		displayGroup.POST("/upload", display.Upload)            //上传图片
		displayGroup.GET("/getall", display.GetAllImages)       //获取所有图片
		displayGroup.GET("/get/:id", display.GetImageByID)      //根据ID获取单张图片
		displayGroup.PUT("/update/:id", display.Update)         //更新图片
		displayGroup.DELETE("/delete/:id", display.DeleteImage) //删除图片
	}
}
