package observer

import (
	"context"
	"fmt"
	"sync"
)

type Event struct {
	Topic string
	Val   interface{}
}

// Observer 要实现 OnChange, 用于向 EventBus 暴漏出通知自己的方式。
// 并且在方法内部实现好当关注对象发生变更时，
// 自己需要采取的处理逻辑。(调用EventBus的Publish方法)

// 观察者抽象类
type Observer interface {
	// Update
	OnChange(ctx context.Context, e *Event) error
}

// 定义抽象EvenBus
type EventBus interface {
	Subscribe(topic string, o Observer)
	Unsubscribe(topic string, o Observer)
	Publish(ctx context.Context, e *Event)
}

// ===== 简单的观察者实现 =======

type BaseObserver struct {
	name string
}

func NewBaseObserver(name string) *BaseObserver {
	return &BaseObserver{
		name: name,
	}
}

func (b *BaseObserver) OnChange(ctx context.Context, e *Event) error {
	fmt.Printf("Observer: %s, event key:%s, event val:%v\n", b.name, e.Topic, e.Val)
	// ...
	return nil
}

// EvenBus 需要实现 Subscribe 和 Unsubscribe 方法暴漏给观察者
// 用于新增或删除订阅关系

type BaseEventBus struct {
	mux       sync.RWMutex
	observers map[string]map[Observer]struct{}
}

func NewBaseEventBus() BaseEventBus {
	return BaseEventBus{
		observers: make(map[string]map[Observer]struct{}),
	}
}

func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	if _, ok := b.observers[topic]; !ok {
		b.observers[topic] = make(map[Observer]struct{})
	}
	b.observers[topic][o] = struct{}{}
}

func (b *BaseEventBus) Unsubscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	delete(b.observers[topic], o)
}

/* ================ */

// 同步模式

// EventBus收到变更事件 Event 时，会根据Event类型 Topic 匹配对应的观察者列表 observers
// 然后用串行遍历分别调用 `Observer.OnChange`方法对每个观察者进行通知
// 对处理流程中遇到的错误进行聚合，放到 handleErr 方法中进行统一的后处理.
type SyncEventBus struct {
	BaseEventBus
}

func NewSyncEventBus() *SyncEventBus {
	return &SyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
}

func (s *SyncEventBus) Publish(ctx context.Context, e *Event) {
	s.mux.RLock()
	subscribers := s.observers[e.Topic]
	s.mux.RUnlock()

	errs := make(map[Observer]error)
	for subscriber := range subscribers {
		if err := subscriber.OnChange(ctx, e); err != nil {
			errs[subscriber] = err
		}
	}

	s.handleErr(ctx, errs)
}

// 简化版 handleErr
// 实际需要针对遇到的错误采取，如重试或告知之类的操作
func (s *SyncEventBus) handleErr(ctx context.Context, errs map[Observer]error) {
	for o, err := range errs {
		// 处理 Publish 失败的 observer
		fmt.Printf("observer: %v, err: %v", o, err)
	}
}
