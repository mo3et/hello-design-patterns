package singleton

import "sync"

var (
	goInstance *GoInstance
	once       sync.Once
)

type GoInstance struct {
	Name string
}

// 使用 once.Do 来保证 某个对象只会初始化一次
//
// 注意: 这个 once.Do 只会被执行一次
// 哪怕 Do func里面发生一次，对象初始化失败了，这个 Do 函数也不会被再次执行了
func GetGoInstance(name string) *GoInstance {
	if goInstance == nil {
		once.Do(func() {
			goInstance = &GoInstance{
				Name: name,
			}
		})
	}
	return goInstance
}
