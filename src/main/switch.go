package main

import "fmt"

func choose() {
	u := user_ya{
		name: "haha",
	}

	// 可比较类型
	switch u {
	case user_ya{
		name: "hahaha",
	}:
		fmt.Printf("yeah this is hahaha")
	default:
		fmt.Printf("default is haha")
	}
}

type user_ya struct {
	name string
}

func main() {
	choose()
}
