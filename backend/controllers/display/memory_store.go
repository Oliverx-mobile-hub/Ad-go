package display

import (
	"fmt"

	upload_image "backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 内存存储层结构体
type MysqlStore struct {
	DB *gorm.DB
}

// 1、NewMysqlStore 初始化数据库连接（open）
func NewMysqlStore() (*MysqlStore, error) {
	dsn := "root:zmkj@1024@Oliver@tcp(172.16.189.83:3306)/Addb?charset=utf8mb4&parseTime=True&loc=Local" //注意Addb为数据库名称
	// 1.1、连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("mysql连接失败:%v", err)
	}
	//1.2、成功连接，自动建表
	if err := db.AutoMigrate(&upload_image.UploadRequest{}); err != nil {
		return nil, fmt.Errorf("mysql建表失败:%v", err)
	}
	//1.3、返回mysqlStore实例
	return &MysqlStore{DB: db}, nil
}

// Upload 上传图片方法
func (s *MysqlStore) Upload(req *upload_image.UploadRequest) error {
	return s.DB.Create(req).Error
}

// update 更新图片方法
func (s *MysqlStore) Update(req *upload_image.UploadRequest) error {
	return s.DB.Save(req).Error
}

// GetImageByID 根据ID获取图片方法
func (s *MysqlStore) GetImageByID(id string) (*upload_image.UploadRequest, error) {
	var image upload_image.UploadRequest
	err := s.DB.Where("id = ?", id).First(&image).Error
	if err != nil {
		return nil, err
	}
	return &image, nil
}

// GetAllImages 获取所有图片方法
func (s *MysqlStore) GetAllImages() ([]upload_image.UploadRequest, error) {
	var images []upload_image.UploadRequest
	err := s.DB.Find(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}

// DeleteImageByID 根据ID删除图片方法
func (s *MysqlStore) DeleteImageByID(id string) error {
	return s.DB.Where("id = ?", id).Delete(&upload_image.UploadRequest{}).Error
}
