package abstractfactory

import "fmt"

// Abstract Factory
type ProductFactory interface {
	CreateTV() TV
	CreatePhone() Phone
}

// TV产品接口
type TV interface {
	Display()
}

// 手机产品接口
type Phone interface {
	Call()
}

// 小米产品工厂
type XiaomiFactory struct{}

// 创建小米电视
func (f XiaomiFactory) CreateTV() TV {
	return XiaomiTV{}
}

// 创建小米手机
func (f XiaomiFactory) CreatePhone() Phone {
	return XiaomiPhone{}
}

// 华为产品工厂
type HuaweiFactory struct{}

// CreateTV 创建华为电视
func (f HuaweiFactory) CreateTV() TV {
	return HuaweiTV{}
}

// CreatePhone 创建华为手机
func (f HuaweiFactory) CreatePhone() Phone {
	return HuaweiPhone{}
}

type XiaomiTV struct{}

func (tv XiaomiTV) Display() {
	fmt.Println("Display in Xiaomi TV.")
}

type XiaomiPhone struct{}

func (phone XiaomiPhone) Call() {
	fmt.Println("Call in XiaomiPhone")
}

type HuaweiTV struct{}

func (tv HuaweiTV) Display() {
	fmt.Println("Display in HuaweiTV")
}

type HuaweiPhone struct{}

func (phone HuaweiPhone) Call() {
	fmt.Println("Call in HuaweiPhone")
}

