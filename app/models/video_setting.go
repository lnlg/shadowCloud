package models

import "shadowCloud/internal/global"

type VideoSetting struct {
	ID        uint      `json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	Notes     string    `json:"notes"`
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
	Deleted   int       `json:"deleted"`
}

// 设置表名
func (VideoSetting) TableName() string {
	return "video_setting"
}

// 添加配置信息
func AddVideoSetting(key string, value string, notes string) (uint, error) {
	//判断key是否存在
	if GetVideoSetting(key) != "" {
		//key已存在，更新value
		rows, err := UpdateVideoSetting(key, value)
		if err != nil {
			return 0, err
		}
		return uint(rows), nil
	} else {
		data := VideoSetting{
			Key:   key,
			Value: value,
			Notes: notes,
		}
		err := global.Db.Create(&data).Error
		return data.ID, err
	}
}

// 根据key获取配置value信息
func GetVideoSetting(key string) string {
	var data VideoSetting
	err := global.Db.Where("`key` = ? AND `deleted` = 0", key).First(&data).Error
	if err != nil {
		return ""
	}
	return data.Value
}

// 更新配置信息
func UpdateVideoSetting(key string, value string) (int64, error) {
	Result := global.Db.Model(&VideoSetting{}).Where("`key` = ?", key).Update("value", value)
	return Result.RowsAffected, Result.Error
}
