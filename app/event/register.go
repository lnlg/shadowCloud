package event

import (
	"shadowCloud/app/event/listener"
	"shadowCloud/internal/event"
	"shadowCloud/internal/global"
)

var listenerList = []event.ListenerInterface{
	listener.FooListener{},
	listener.BarListener{},
}

// 注册事件
func RegisterAppEvent() {
	for _, listener := range listenerList {
		global.Event.AddListener(listener)
	}
}
