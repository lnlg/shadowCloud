package listener

import (
	"fmt"
	appEvent "shadowCloud/app/event/event"
	"shadowCloud/internal/event"
)

type FooListener struct{}

// 监听事件
func (f FooListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.FooEvent{},
	}
}

// 处理事件
func (f FooListener) Process(e event.EventInterface) {
	fmt.Println("foo listener process event:", e, e.Name())
	// println(e.(*appEvent.FooEvent).Id)
	// println(e.(*appEvent.FooEvent).Username)
}
