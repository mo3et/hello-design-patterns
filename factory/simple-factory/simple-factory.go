package simplefactory

import "fmt"

// Simple Factory

type Pet interface {
	call()
}

type Dog struct{}

func (*Dog) call() {
	fmt.Println("汪汪汪——————")
}

type Cat struct{}

func (*Cat) call() {
	fmt.Println("喵喵喵————")
}

// 宠物工厂类
type PetFactory struct{}

// 根据传入类型创建具体产品
func (*PetFactory) CreatePet(like string) Pet {
	if like == "cat" {
		return &Cat{}
	} else if like == "dog" {
		return &Dog{}
	}
	return nil
}
