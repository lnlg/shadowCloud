package models

import "shadowCloud/internal/global"

type VideoDownloadUrl struct {
	ID        uint      `json:"id"`
	Url       string    `json:"url"`
	HashUrl   string    `json:"hash_url"`
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
	Deleted   int       `json:"deleted"`
}

// 设置表名
func (VideoDownloadUrl) TableName() string {
	return "video_download_url"
}

// 添加视频下载链接
func AddVideoDownloadUrl(url string, hashUrl string) (uint, error) {
	data := VideoDownloadUrl{
		Url:     url,
		HashUrl: hashUrl,
	}
	err := global.Db.Create(&data).Error
	return data.ID, err
}

// 获取下载地址
func GetVideoDownloadUrl(hashUrl string) (VideoDownloadUrl, error) {
	var data VideoDownloadUrl
	err := global.Db.Where("hash_url = ? AND deleted = 0", hashUrl).First(&data).Error
	return data, err
}
