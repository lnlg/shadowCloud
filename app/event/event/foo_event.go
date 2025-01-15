package event

// 定义事件实体
type TestEvent struct {
	Id       int
	Username string
}

// 事件名称
func (f *TestEvent) Name() string {
	return "test_event"
}
