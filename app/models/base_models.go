package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 1.自定义时间格式
type LocalTime time.Time

// 2.实现 MarshalJSON 方法实现数据解析
func (t LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02 15:04:05") + `"`), nil
}

// 3.实现 Scan 方法实现数据解析
func (t *LocalTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*t = LocalTime(v)
		return nil
	case string:
		layout := "2006-01-02 15:04:05"
		theTime, err := time.Parse(layout, value.(string))
		if err != nil {
			return err
		}
		*t = LocalTime(theTime)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// 4.实现 Value 方法实现数据解析
func (t LocalTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// 5. 自定义 BaseModel，结构和 gorm.Model 一致，将 time.Time 替换为 LocalTime
type BaseModel struct {
	ID        int64     `json:"id"`
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
	IsDeleted int       `json:"is_deleted"`
}
