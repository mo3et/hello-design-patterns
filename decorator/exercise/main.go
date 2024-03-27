package main

import "fmt"

// 装饰模式 咖啡加糖
// https://kamacoder.com/problempage.php?pid=1086

// 咖啡接口
type Coffee interface {
	brew()
}

// 具体的黑咖啡类
type BlackCoffee struct{}

func (bc *BlackCoffee) brew() {
	fmt.Println("Brewing Black Coffee")
}

// 具体的拿铁类
type Latte struct{}

func (l *Latte) brew() {
	fmt.Println("Brewing Latte")
}

// 装饰者抽象类
type Decorator struct {
	coffee Coffee
}

func (d *Decorator) brew() {
	d.coffee.brew()
}

// Details Milk Decorator
type MilkDecorator struct {
	Decorator
}

func (md *MilkDecorator) brew() {
	md.Decorator.brew()
	fmt.Println("Adding Milk")
}

// 具体的糖装饰者类
type SugarDecorator struct {
	Decorator
}

func (sd *SugarDecorator) brew() {
	sd.Decorator.brew()
	fmt.Println("Adding Sugar")
}

func main() {
	for {
		var coffeeType, condimentType int
		if _, err := fmt.Scan(&coffeeType, &condimentType); err != nil {
			break
		}

		// 根据输入制作咖啡
		var coffee Coffee
		if coffeeType == 1 {
			coffee = &BlackCoffee{}
		} else if coffeeType == 2 {
			coffee = &Latte{}
		} else {
			fmt.Println("Invalid coffee type")
			continue
		}

		// 根据输入添加调料
		if condimentType == 1 {
			coffee = &MilkDecorator{Decorator: Decorator{coffee: coffee}}
		} else if condimentType == 2 {
			coffee = &SugarDecorator{Decorator: Decorator{coffee: coffee}}
		} else {
			fmt.Println("Invalid condiment type")
			continue
		}

		// 输出制作过程
		coffee.brew()

	}
}
