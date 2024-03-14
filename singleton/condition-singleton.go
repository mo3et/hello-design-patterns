package singleton

import "sync"

// Ref: https://juejin.cn/post/7124720007447052302

type Singleton struct {
	Name string
}

var (
	lock      sync.RWMutex
	instances = map[string]*Singleton{}
)

// sync.Once 是仅执行一次。
//
// 这里需要判断map中对应的value是否存在
// Once就要初始化多次实例，所以不适合。
func GetInstance(key string) *Singleton {
	lock.RLock()
	// 需要判断map中对应key 是否存在
	if value, ok := instances[key]; ok {
		lock.RUnlock()
		return value
	}

	lock.RUnlock()

	lock.Lock()
	defer lock.Unlock()
	if value, ok := instances[key]; ok {
		return value
	}
	instance := &Singleton{Name: "zhangsan"}
	instances[key] = instance
	return instance
}
