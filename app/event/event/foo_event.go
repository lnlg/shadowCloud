package event

// 定义事件实体
type FooEvent struct {
	Id       int
	Username string
}

// 事件名称
func (f *FooEvent) Name() string {
	return "foo_event"
}
