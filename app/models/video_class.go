package models

import "shadowCloud/internal/global"

type VideoClass struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Sort      int       `json:"sort"`
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
	Deleted   int       `json:"deleted"`
}

// TableName 表名
func (VideoClass) TableName() string {
	return "video_class"
}

// 添加视频分类
func AddVideoClass(name string, sort int) (uint, error) {
	data := VideoClass{
		Name: name,
		Sort: sort,
	}
	err := global.Db.Create(&data).Error
	return data.ID, err
}

// 获取视频分类列表
func GetVideoClassList() ([]VideoClass, error) {
	var list []VideoClass
	err := global.Db.Where("deleted = 0").Order("`sort` asc").Find(&list).Error
	return list, err
}
