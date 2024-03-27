package main

import "testing"

func Test_Builder(t *testing.T) {
	// 创建 Food 建造者实例
	fb := NewFoodBuilder()
	// 通过链式调用完成属性设置与实例建造
	food1, err := fb.Type("苹果").Cost(12.12).Brand("山东红富士").Build()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("food: %+v", food1)
	}

	// 通过链式调用完成属性设置与实例建造
	food2, err := fb.Type("芒果").Name("我是大芒果1号").Cost(30.30).Build()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("food: %+v", food2)
	}
}
