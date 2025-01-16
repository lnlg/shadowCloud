package models

import (
	"shadowCloud/internal/global"
)

type VideoList struct {
	Id             uint      `json:"id"`
	TypeId         int       `json:"type_id"`
	VideoName      string    `json:"video_name"`
	VideoUrl       string    `json:"video_url"`
	ImageUrl       string    `json:"image_url"`
	OrgDownloadUrl string    `json:"org_download_url"`
	DownloadState  int       `json:"download_state"`
	Md5File        string    `json:"md5_file"`
	CreatedAt      LocalTime `json:"created_at"`
	UpdatedAt      LocalTime `json:"updated_at"`
	Deleted        int64     `json:"deleted"`
}

// 设置表名
func (VideoList) TableName() string {
	return "video_list"
}

// 添加视频
func AddVideo(video *VideoList) (uint, error) {
	err := global.Db.Create(video).Error
	return video.Id, err
}

// 获取视频列表-分页 每次获取指定条数
func GetVideoList(type_id int, page int, pageSize int) ([]VideoList, error) {
	var videos []VideoList
	err := global.Db.Where("`type_id` = ? AND `deleted` = 0", type_id).Order("`id` DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&videos).Error
	return videos, err
}

// 根据type_id 获取总条数
func GetVideoListCount(type_id int) (int64, error) {
	var count int64
	err := global.Db.Model(&VideoList{}).Where("`type_id` = ? AND `deleted` = 0", type_id).Count(&count).Error
	return count, err
}

// 根据id 获取视频
func GetVideoById(id uint) (VideoList, error) {
	var video VideoList
	err := global.Db.Where("`id` = ? AND `deleted` = 0", id).First(&video).Error
	return video, err
}

// 根据id 更新download_state
func UpdateVideoDownloadState(id uint, downloadState int) (int64, error) {
	res := global.Db.Model(&VideoList{}).Where("`id` = ?", id).Update("download_state", downloadState)
	return res.RowsAffected, res.Error
}

// 获取随机未下载的视频信息
func GetRandomVideoInfo() (VideoList, error) {
	var video VideoList
	err := global.Db.Where("`download_state` = 0 AND `deleted` = 0").Order("RAND()").First(&video).Error
	return video, err
}
