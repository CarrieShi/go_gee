package main

import "fmt"

func main() {
	arr := []int{10, 11, 12}

	for index, value := range arr {
		fmt.Printf("%d ==> %d \n", index, value)
	}

	// 匿名函数 _ 代替 index
	for _, value := range arr {
		fmt.Printf("NAN ==> %d \n", value)
	}

	// only index 1
	for index := range arr {
		fmt.Printf("%d ==> NAN \n", index)
	}

	// only index 2
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ==> NAN \n", i)
	}
}
