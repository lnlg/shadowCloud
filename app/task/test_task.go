package task

import (
	"fmt"
)

// TestTask is a task that prints "foo task executed" every 3 seconds.
type TestTask struct{}

// https://godoc.org/github.com/robfig/cron
// 每半个小时执行一次  "0 30 * * * *"
// 每小时执行一次     "@hourly"
// 每小时三十分钟执行一次 "@every 1h30m"
// 每3秒执行一次       "@every 3s"
// CRON表达式格式：
// Field name   | Mandatory? | Allowed values  | Allowed special characters
// ----------   | ---------- | --------------  | --------------------------
// Seconds      | Yes        | 0-59            | * / , -
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

// 定义任务循环周期
func (f *TestTask) Spec() string {
	return "@every 30s"
}

// 定义任务执行函数
func (f *TestTask) Fn() func() {
	return func() {
		fmt.Println("test_task 开始执行!")
	}
}
