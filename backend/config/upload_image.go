package config

// 1、上传图片结构体
type UploadRequest struct {
	ID       string `gorm:"primaryKey;size:64" json:"id"` // 图片ID 主键
	ImageURL string `json:"image_url"`                    // 图片URL
	Desc     string `json:"desc"`                         // 图片描述
}
