package listener

import (
	"fmt"
	appEvent "shadowCloud/app/event/event"
	"shadowCloud/internal/event"
)

type BarListener struct{}

// 监听事件
func (b BarListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.FooEvent{},
	}
}

// 处理事件
func (b BarListener) Process(e event.EventInterface) {
	fmt.Println("bar listener process event:", e, e.Name())
}
