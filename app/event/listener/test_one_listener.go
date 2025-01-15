package listener

import (
	"fmt"
	appEvent "shadowCloud/app/event/event"
	"shadowCloud/internal/event"
)

type TestOneListener struct{}

// 监听事件
func (t TestOneListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.TestEvent{},
	}
}

// 处理事件
func (t TestOneListener) Process(e event.EventInterface) {
	fmt.Println("foo listener process event:", e, e.Name())
	// println(e.(*appEvent.FooEvent).Id)
	// println(e.(*appEvent.FooEvent).Username)
}
