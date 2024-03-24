package observer

import (
	"context"
	"testing"
)

func Test_baseEventBus(t *testing.T) {
	observerA := NewBaseObserver("a")
	observerB := NewBaseObserver("b")
	observerC := NewBaseObserver("c")
	observerD := NewBaseObserver("d")

	bbus := NewSyncEventBus()
	topic := "mission complete"
	bbus.Subscribe(topic, observerA)
	bbus.Subscribe(topic, observerB)
	bbus.Subscribe(topic, observerC)
	bbus.Subscribe(topic, observerD)

	bbus.Publish(context.Background(), &Event{
		Topic: topic,
		Val:   "order_id: xxx",
	})
}
