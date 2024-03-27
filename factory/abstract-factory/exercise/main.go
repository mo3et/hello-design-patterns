package main

import "fmt"

// 抽象椅子接口
type Chair interface {
	showInfo()
}

// 具体现代风格椅子
type ModernChair struct{}

func (mc *ModernChair) showInfo() {
	fmt.Println("modern chair")
}

// 具体古典风格椅子
type ClassicalChair struct{}

func (cc *ClassicalChair) showInfo() {
	fmt.Println("classical chair")
}

// 抽象沙发接口
type Sofa interface {
	displayInfo()
}

// 具体现代风格沙发
type ModernSofa struct{}

func (ms *ModernSofa) displayInfo() {
	fmt.Println("modern sofa")
}

// 具体古典风格沙发
type ClassicalSofa struct{}

func (cs *ClassicalSofa) displayInfo() {
	fmt.Println("classical sofa")
}

// 抽象家居工厂接口
type FurnitureFactory interface {
	createChair() Chair
	createSofa() Sofa
}

// 具体现代风格家居工厂
type ModernFurnitureFactory struct{}

func (mf *ModernFurnitureFactory) createChair() Chair {
	return &ModernChair{}
}

func (mf *ModernFurnitureFactory) createSofa() Sofa {
	return &ModernSofa{}
}

// 具体古典风格家居工厂
type ClassicalFurnitureFactory struct{}

func (cf *ClassicalFurnitureFactory) createChair() Chair {
	return &ClassicalChair{}
}

func (cf *ClassicalFurnitureFactory) createSofa() Sofa {
	return &ClassicalSofa{}
}

func main() {
	// 读取订单数量
	var N int
	fmt.Scan(&N)

	// 处理每个订单
	for i := 0; i < N; i++ {
		// 读取家具类型
		var furnitureType string
		fmt.Scan(&furnitureType)

		// 创建相应风格的家居装饰品工厂
		var factory FurnitureFactory
		if furnitureType == "modern" {
			factory = &ModernFurnitureFactory{}
		} else if furnitureType == "classical" {
			factory = &ClassicalFurnitureFactory{}
		}

		// 根据工厂生产椅子和沙发
		chair := factory.createChair()
		sofa := factory.createSofa()

		// 输出家具信息
		chair.showInfo()
		sofa.displayInfo()
	}
}
