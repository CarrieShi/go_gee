package main

import "fmt"

func main() {
	// 初始化 长度 内容
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1 %v len %d cap %d \n", a1, len(a1), cap(a1))
	fmt.Printf("a1[1] %#v \n", a1[1])

	// 初始化 长度 内容默认为零值
	//a2 := [4]int{}
	var a2 [4]int
	//var a2 = [4]int{0, 0, 0, 0}
	fmt.Printf("a2 %v len %d cap %d \n", a2, len(a2), cap(a2))
	fmt.Printf("a2[1] %#v \n", a2[1])
}
