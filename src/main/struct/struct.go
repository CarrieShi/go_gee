package main

import "fmt"

// 没有构造函数，如何创建实例

func main() {
	// *ToyDuck ==> 取地址，指针
	duck1 := &ToyDuck{}
	duck1.Swim("No1")

	// 结构体 实例
	duck2 := ToyDuck{}
	duck2.Swim("No2")

	// *ToyDuck 指针 分配好内存，内存初始化（比特位 置为零值）
	// 构造函数 执行里面的逻辑，goland 没有、不会、no
	duck3 := new(ToyDuck)
	duck3.Swim("No3")

	// 声明即分配好内存，有零值
	var duck4 ToyDuck
	duck4.Swim("No4")

	// 指针
	//var duck5 *ToyDuck
	//panic: runtime error: invalid memory address or nil pointer dereference
	//duck5.Swim("No5")

	duck6 := ToyDuck{
		Color: "yellow",
		Price: 10,
	}
	duck6.Color = "red"
	duck6.Swim("No6")
}

type ToyDuck struct {
	Color string
	Price uint64
}

func (duck ToyDuck) Swim(no string) {
	fmt.Printf("玩具鸭子 " + no + " 游泳\n")
}
