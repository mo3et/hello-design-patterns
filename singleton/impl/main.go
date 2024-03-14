package main

import (
	"fmt"
	"sync"
)

type ShoppingCart struct {
	products map[string]int
}

// 单例实例以及同步

var (
	instance *ShoppingCart
	once     sync.Once
)

// GetInstance

func GetInstance() *ShoppingCart {
	once.Do(func() {
		instance = &ShoppingCart{products: make(map[string]int)}
	})
	return instance
}

// 添加产品到购物车
func (cart *ShoppingCart) AddProduct(name string, quantity int) {
	cart.products[name] += quantity
}

// 打印购物清单
func (cart *ShoppingCart) PrintInvoice() {
	for name, quantity := range cart.products {
		fmt.Printf("%s %d\n", name, quantity)
	}
}

func main() {
	var name string
	var quantity int

	cart := GetInstance()

	for {
		// 从标准输入读取商品名称和数量
		if _, err := fmt.Scanf("%s %d", &name, &quantity); err != nil {
			break // 当输入结束或发生错误时停止读取
		}
		cart.AddProduct(name, quantity)
	}
	cart.PrintInvoice()
}
