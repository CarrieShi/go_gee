package main

import (
	"fmt"
	"time"
)

func main() {

	// 闭包 = 匿名函数 + 定义上下文
	name := "TOM"
	sayHiTo := func() {
		fmt.Printf("say hi to %s \n", name)
	}
	sayHiTo()

	fmt.Println(ReturnClosure("JIM")())

	Delay()
	time.Sleep(time.Second)
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "hello, " + name
	}
}

func Delay() {
	// 延迟绑定
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("hello, this is %d \n", i)
		})
	}

	// 执行匿名函数
	// 结果全是 hello, this is 10 ==> 最近一次的 i
	// 而不是
	// hello, this is 1
	// hello, this is ...
	for _, fn := range fns {
		fn()
	}
}
