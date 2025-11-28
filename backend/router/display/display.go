package display

import (
	"backend/controllers/display"

	"github.com/gin-gonic/gin"
)

// 公共路由 - 不需要认证
func RegisterPublicDisplayRoutes(g *gin.RouterGroup) {
	publicGroup := g.Group("/display")
	{
		// 公共接口：获取图片信息，不需要认证
		publicGroup.GET("/get/:id", display.GetImageByID) //根据ID获取单张图片
	}
}

// 受保护路由 - 需要认证
func RegisterProtectedDisplayRoutes(g *gin.RouterGroup) {
	protectedGroup := g.Group("/display")
	{
		// 受保护接口：需要认证才能访问
		protectedGroup.POST("/upload", display.Upload)            //上传图片
		protectedGroup.GET("/getall", display.GetAllImages)       //获取所有图片
		protectedGroup.PUT("/update/:id", display.Update)         //更新图片
		protectedGroup.DELETE("/delete/:id", display.DeleteImage) //删除图片
	}
}
