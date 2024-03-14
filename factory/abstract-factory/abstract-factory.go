package abstractfactory

import (
	"fmt"
)

// Abstract Factory

// 抽象工厂模式接口
type Company interface {
	slogan()
}

// 小米
type Xiaomi struct{}

func (Xiaomi) slogan() {
	fmt.Println("Are you OK")
}

// 哔哩哔哩
type Bilibili struct{}

func (Bilibili) slogan() {
	fmt.Println("Bilibili 干杯")
}

// Factory 工厂接口
type CompanyShowor interface {
	ShowCompany(name string) Company
}

// 中国工厂
type ChinaComFactory struct{}

func (ChinaComFactory) ShowCompany(name string) Company {
	if name == "xiaomi" {
		return &Xiaomi{}
	} else if name == "bilibili" {
		return &Bilibili{}
	}
	return nil
}

// 美国工厂
type UsComFactory struct{}

func (UsComFactory) ShowCompany(name string) Company {
	if name == "apple" {
		return &Apple{}
	} else if name == "samsung" {
		return &Samsung{}
	}
	return nil
}

// Apple
type Apple struct{}

func (Apple) slogan() {
	fmt.Println("Only apple can do!")
}

// Samsung
type Samsung struct{}

func (Samsung) slogan() {
	fmt.Println("Boooooom!!")
}

// 工厂提供者(实现)
type CompanyFactoryStore struct {
	shower CompanyShowor
}

func (show *CompanyFactoryStore) showCompany(name string) Company {
	return show.shower.ShowCompany(name)
}
