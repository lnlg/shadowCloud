package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 1.自定义时间格式
type LocalTime time.Time

// 2.实现 MarshalJSON 方法
func (t LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02 15:04:05") + `"`), nil
}

// 3.实现 Scan 方法
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

// 4.实现 Value 方法
func (t LocalTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}
