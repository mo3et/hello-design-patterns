package main

import (
	"fmt"
)

/*
2. 工厂方法模式 积木工厂
https://kamacoder.com/problempage.php?pid=1076
*/

// 抽象积木接口
type Block interface {
	produce()
}

// 具体圆形积木实现
type CircleBlock struct{}

func (c *CircleBlock) produce() {
	fmt.Println("Circle Block")
}

// 具体方形积木实现
type SquareBlock struct{}

func (s *SquareBlock) produce() {
	fmt.Println("Square Block")
}

// 抽象积木工厂接口
type BlockFactory interface {
	createBlock() Block
}

// 具体圆形积木工厂实现
type CircleBlockFactory struct{}

func (cf *CircleBlockFactory) createBlock() Block {
	return &CircleBlock{}
}

// 具体方形积木工厂实现
type SquareBlockFactory struct{}

func (sf *SquareBlockFactory) createBlock() Block {
	return &SquareBlock{}
}

// 积木工厂系统
type BlockFactorySystem struct {
	blocks []Block
}

func (bfs *BlockFactorySystem) produceBlocks(factory BlockFactory, quantity int) {
	for i := 0; i < quantity; i++ {
		block := factory.createBlock()
		bfs.blocks = append(bfs.blocks, block)
		block.produce()
	}
}

func (bfs *BlockFactorySystem) getBlocks() []Block {
	return bfs.blocks
}

func main() {
	// 创建积木工厂系统
	factorySystem := &BlockFactorySystem{}

	// 读取生产次数
	var productionCount int
	fmt.Scan(&productionCount)

	// 读取每次生产的积木类型和数量
	for i := 0; i < productionCount; i++ {
		var blockType string
		var quantity int
		fmt.Scan(&blockType, &quantity)

		if blockType == "Circle" {
			factorySystem.produceBlocks(&CircleBlockFactory{}, quantity)
		} else if blockType == "Square" {
			factorySystem.produceBlocks(&SquareBlockFactory{}, quantity)
		}
	}
}
