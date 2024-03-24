package main

import "testing"

func Test_decorator(t *testing.T) {
	// 一碗干净的米饭
	rice := NewRice()
	rice.Eat()

	// 一碗干净的面条
	noodle := NewNoodle()
	noodle.Eat()

	// 米饭加个煎蛋
	rice = NewFriedEggDecorator(rice)
	rice.Eat()

	// 面条加份火腿
	noodle = NewHamSausageDecorator(noodle)
	noodle.Eat()

	// 米饭再分别加个煎蛋和一份老干妈
	rice = NewFriedEggDecorator(rice)
	rice = NewLaoGanMaDecorator(rice)
	rice.Eat()
}
