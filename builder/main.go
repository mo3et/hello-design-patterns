package main

import (
	"errors"
	"log"
)

// 小徐先生 [Builder pattern](https://mp.weixin.qq.com/s?__biz=MzkxMjQzMjA0OQ==&mid=2247484474&idx=1&sn=739d0e48fa735a3d2bb3328db0d35d08&chksm=c10c4ae4f67bc3f210540b371943a30613ae7e290010a93660b86d427a77380dac7324ca6bd9&scene=178&cur_album_id=2935694957926088707#rd)

type Food struct {
	typ    string
	name   string
	weight float64
	brand  string
	cost   float64
}

func NewFood(typ string, name string, weight float64, brand string, cost float64) *Food {
	return &Food{
		typ:    typ,
		name:   name,
		weight: weight,
		brand:  brand,
		cost:   cost,
	}
}

type FoodBuilder struct {
	Food
}

// 定义 FoodBuilder 类
// 通过 embed Food 方式，使其包含有 Food 的所有成员属性。
func NewFoodBuilder() *FoodBuilder {
	return &FoodBuilder{}
}

// 为 FoodBuilder 声明好通过链式调用完成 Food 成员属性设值的方法
func (f *FoodBuilder) Type(typ string) *FoodBuilder {
	f.typ = typ
	return f
}

func (f *FoodBuilder) Name(name string) *FoodBuilder {
	f.name = name
	return f
}

func (f *FoodBuilder) Weight(weight float64) *FoodBuilder {
	f.weight = weight
	return f
}

func (f *FoodBuilder) Brand(brand string) *FoodBuilder {
	f.brand = brand
	return f
}

func (f *FoodBuilder) Cost(cost float64) *FoodBuilder {
	f.cost = cost
	return f
}

// Build方法，可以基于用户预设好的成员属性，完成Food实例的构造
// 在此，还可以利用 Build方法产生的切面，来对 Food 类中所要求的字段进行校验
// 包括 typ 和 name 两个字段.

func (f *FoodBuilder) Build() (*Food, error) {
	if f.typ == "" {
		return nil, errors.New("miss type info")
	}
	if f.name == "" {
		return nil, errors.New("miss name info")
	}

	return &Food{
		typ:    f.typ,
		name:   f.name,
		brand:  f.brand,
		weight: f.weight,
		cost:   f.cost,
	}, nil
}

func main() {
	// 创建 Food 建造者实例
	fb := NewFoodBuilder()
	// 通过链式调用完成属性设置与实例建造
	food1, err := fb.Type("苹果").Cost(12.12).Brand("山东红富士").Build()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("food: %+v", food1)
	}

	// 通过链式调用完成属性设置与实例建造
	food2, err := fb.Type("芒果").Name("我是大芒果1号").Cost(30.30).Build()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("food: %+v", food2)
	}
}
