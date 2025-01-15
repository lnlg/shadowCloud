package event

// 监听器接口
type ListenerInterface interface {
	Listen() []EventInterface
	Process(EventInterface)
}
