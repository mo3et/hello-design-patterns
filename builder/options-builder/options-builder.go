package main

import (
	"fmt"
)

// Options 模式

/*
如果有一个BigClass 类，里面有很多成员属性
其中相当一部分是选填的，因此构造函数的入参组合也不胜枚举
*/

/* type BigClass struct {
    name   string
    age    int
    sex    string
    weight float64
    height float64
    width  float64
    fieldA string
    fieldB string
    fieldC string
    // ...
} */

// 我们对BigClass 进行改造，将构造过程需要关心的成员属性都聚合在配置项类 Options中，让 BigClass 直接  embed Options

type BigClass struct {
	Options
}

type Options struct {
	name   string
	age    int
	sex    string
	weight float64
	height float64
	width  float64
	fieldA string
	fieldB string
	fieldC string
}

// 定义一个配置函数类 Options，对应类型是 func(*Options)
// 通过在入参中接口收到 Options 配置项类的指针
// 使Option 在运行过程中能够完成对 Options 当中成员属性的赋值修改

/*
!!! NOTICE:
Options 是聚合成员属性的配置项类
另一个 Option 是*配置函数*的类型
目的是为了通过Option的运行完成对 Options 中的成员属性赋值
*/

// 是配置函数的type，为了完成对Options中成员属性的赋值
// 这个函数 例如`WithValue(20)` 可以接受后面的*Options地址 作为闭包的入参

type OptionFunc func(opts *Options)

// type Option func(opts *Options)

// 下面根据需要的成员属性范围，提前声明好一系列配置器方法
// 统一以WithXXX 进行命名，在入参传入接收用户设置的成员属性值，出参为 Option配置函数的类型
// 通过闭包，将用户传入的属性值赋到 Options 配置类的成员属性当中。

func WithName(name string) OptionFunc {
	return func(opts *Options) {
		opts.name = name
	}
}

func WithAge(age int) OptionFunc {
	return func(opts *Options) {
		opts.age = age
	}
}

func WithSex(sex string) OptionFunc {
	return func(opts *Options) {
		opts.sex = sex
	}
}

func WithWeight(weight float64) OptionFunc {
	return func(opts *Options) {
		opts.weight = weight
	}
}

func WithHeight(height float64) OptionFunc {
	return func(opts *Options) {
		opts.height = height
	}
}

func WithWidth(width float64) OptionFunc {
	return func(opts *Options) {
		opts.width = width
	}
}

func WithFieldA(fieldA string) OptionFunc {
	return func(opts *Options) {
		opts.fieldA = fieldA
	}
}

func WithFieldB(fieldB string) OptionFunc {
	return func(opts *Options) {
		opts.fieldB = fieldB
	}
}

func WithFieldC(fieldC string) OptionFunc {
	return func(opts *Options) {
		opts.fieldC = fieldC
	}
}

// 最后通过一个兜底修复方法 repair, 完成构造 BigClass 实例过程中的一些缺省值设置

// 如没有通过 Option 显式设置字段，则通过repair 方法设置默认值
func repair(opts *Options) {
	if opts.name == "" {
		opts.name = "小明"
	}
	if opts.age == 0 {
		opts.age = 20
	}
}

// 构造器函数 BigClass 的定义，入参用可变长度的 Option list 传入
// 在方法执行中对list进行遍历，依次执行每个 Option 对其成员属性的赋值操作，最后repair保证缺省值的设置。

func NewBigClass(opts ...OptionFunc) *BigClass {
	bigClass := BigClass{}
	for _, opt := range opts {
		opt(&bigClass.Options)
	}

	repair(&bigClass.Options)

	return &bigClass
}

func main() {
	bc1 := NewBigClass(WithAge(10), WithName("Jason"), WithSex("man"))
	bC := BigClass{}
	WithAge(10)(&bC.Options)
	fmt.Println(bc1)
}
