package event

import (
	"shadowCloud/app/event/listener"
	"shadowCloud/internal/event"
	"shadowCloud/internal/global"
)

var listenerList = []event.ListenerInterface{
	listener.TestOneListener{},
	listener.TestTwoListener{},
}

// 注册事件
func RegisterAppEvent() {
	for _, listener := range listenerList {
		global.Event.AddListener(listener)
	}
}
