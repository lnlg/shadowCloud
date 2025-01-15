package listener

import (
	"fmt"
	appEvent "shadowCloud/app/event/event"
	"shadowCloud/internal/event"
)

type TestTwoListener struct{}

// 监听事件
func (t TestTwoListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.TestEvent{},
	}
}

// 处理事件
func (t TestTwoListener) Process(e event.EventInterface) {
	fmt.Println("test_two listener process event:", e, e.Name())
}
