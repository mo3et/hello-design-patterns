package singleton

import "sync"

var (
	instance *Instance
	mtlock   sync.Mutex
)

type Instance struct {
	Name string
}

func SimpleGetInstance(name string) *Instance {
	if instance == nil {
		mtlock.Lock()
		defer mtlock.Unlock()
		if instance == nil {
			// 构造函数,私有的
			instance = &Instance{Name: name}
		}
	}
	return instance
}
