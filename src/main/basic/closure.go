package main

import "fmt"

func main() {

	// 闭包 = 匿名函数 + 定义上下文
	name := "TOM"
	sayHiTo := func() {
		fmt.Printf("say hi to %s \n", name)
	}
	sayHiTo()
}
