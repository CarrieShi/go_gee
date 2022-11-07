package main

import "fmt"

func main() {
	// 输出，堆栈形式，先进后出
	defer func() {
		fmt.Println("aaa")
	}()
	defer func() {
		fmt.Println("bbb")
	}()
	defer func() {
		fmt.Println("ccc")
	}()
}
