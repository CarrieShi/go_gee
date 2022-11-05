package main

import "fmt"

type Swimming interface {
	Swim()
}

type Flying interface {
	Fly()
}

// Duck 使用了组合
type Duck interface {
	Swimming
	Flying
}

type Base struct {
	Name string
}

// Concrete1 使用结构体进行组合
type Concrete1 struct {
	Base
}

// Concrete2 使用结构体指针进行组合
type Concrete2 struct {
	*Base
}

func (b *Base) SayHello() {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}

func (b *Base) SayGoodbye() {
	fmt.Printf("See you , I am base my name is :%s \n", b.Name)
}

func (c Concrete1) SayHello() {
	// c.Name 直接访问了 Base 的 Name 字段
	fmt.Printf("I am base call by c and my name is: %s \n", c.Name)
	// also
	fmt.Printf("I am base call by c and my name is: %s \n", c.Base.Name)

	// 调用了被组合的
	c.Base.SayHello()
	c.SayGoodbye()
}

type Parent struct {
}

type Son struct {
	Parent
}

func (p Parent) SayHi() {
	fmt.Printf("do not call me %s \n", p.Name())
}

func (p Parent) Name() string {
	return "old man"
}

//func (s Son) SayHi() {
//	fmt.Printf("do not call me %s \n", s.Name())
//}

//func (s Son) Name() string {
//	return "kid"
//}

func main() {
	concrete1 := Concrete1{
		Base{
			Name: "gaga",
		},
	}

	concrete1.SayHello()

	son := Son{
		Parent{},
	}

	// papa or kid
	// if son has overwritten say hi,then show kid
	// otherwise, go is different from usual，show old man instead
	son.SayHi()
	// if son has a name, then show kid
	// otherwise, show old man
	fmt.Printf("return %s \n", son.Name())
}
