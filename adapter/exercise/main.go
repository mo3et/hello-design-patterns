package main

import "fmt"

/*
适配器模式 6. 扩展坞
https://kamacoder.com/problempage.php?pid=1085
*/

// USB 接口
type USB interface {
	charge()
}

// TypeC 接口
type TypeC interface {
	chargeWithTypeC()
}

// 适配器类
type TypeCAdapter struct {
	typeC TypeC
}

func (tca *TypeCAdapter) charge() {
	tca.typeC.chargeWithTypeC()
}

// 新电脑类，使用 TypeC 接口
type NewComputer struct{}

func (nc *NewComputer) chargeWithTypeC() {
	fmt.Println("TypeC")
}

// 适配器充电器类，使用 USB 接口
type AdapterCharger struct{}

func (ac *AdapterCharger) charge() {
	fmt.Println("USB Adapter")
}

func main() {
	var N int
	fmt.Scan(&N) // 读取连接次数

	for i := 0; i < N; i++ {
		var choice int
		fmt.Scan(&choice) // 读取用户选择

		// 根据用户的选择创建相应对象
		if choice == 1 {
			newComputer := &NewComputer{}
			adapter := &TypeCAdapter{typeC: newComputer}
			adapter.charge()
		} else if choice == 2 {
			usbAdapter := &AdapterCharger{}
			usbAdapter.charge()
		}
	}
}
