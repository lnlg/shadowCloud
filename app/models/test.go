package models

import "shadowCloud/internal/global"

// Test 模型
type Test struct {
	Id        int64     `json:"id"` // `json:"id"` 接口返回使用小写
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
	IsDeleted int       `json:"is_deleted"`
}

// TableName 表名
func (Test) TableName() string {
	return "test"
}

// 获取一条数据
func GetTestOne(id int64) (Test, error) {
	var test Test
	err := global.Db.First(&test, id).Error
	return test, err
}
