package event

import "sync"

// 事件调度器
type Dispatcher struct {
	listenerMap map[string][]ListenerInterface
	lock        sync.RWMutex
}

// 创建事件调度器
func New() *Dispatcher {
	return &Dispatcher{
		listenerMap: make(map[string][]ListenerInterface),
		lock:        sync.RWMutex{},
	}
}

// 添加监听器
func (d *Dispatcher) AddListener(listener ListenerInterface) {
	if listener == nil {
		return
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	// 获取监听者监听的事件列表
	eventList := listener.Listen()

	// 记录事件和监听者的关系
	for _, event := range eventList {
		d.listenerMap[event.Name()] = append(d.listenerMap[event.Name()], listener)
	}
}

// 触发事件
func (d *Dispatcher) Dispatch(event EventInterface) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	// 获取监听器
	listenersList, exist := d.listenerMap[event.Name()]
	if !exist {
		return
	}
	// 遍历监听器
	for _, listener := range listenersList {
		listener.Process(event)
	}
}

// 移除监听器
func (d *Dispatcher) RemoveListener(listener ListenerInterface) {
	d.lock.Lock()
	defer d.lock.Unlock()
	for _, event := range listener.Listen() {
		delete(d.listenerMap, event.Name())
	}
}
