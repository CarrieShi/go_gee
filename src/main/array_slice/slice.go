package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s1 %v len %d cap %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s1[1] %#v \n", s1[1])

	s2 := make([]int, 3, 5)
	fmt.Printf("s2 %v len %d cap %d \n", s2, len(s2), cap(s2))
	fmt.Printf("s2[1] %#v \n", s2[1])
	//fmt.Printf("s2[4] %#v \n", s2[4]) // panic: runtime error: index out of range [4] with length 3

	s2 = append(s2, 7, 8)
	fmt.Printf("s2[4] %#v \n", s2[4])

	// 扩容 cap < 1000 , 扩容后 = cap * 2
	s2 = append(s2, 9)
	fmt.Printf("s2 %v len %d cap %d \n", s2, len(s2), cap(s2))

	// cap > 1000 扩容 2048, 规律？
	s3 := make([]int, 1000, 1000)
	s3 = append(s3, 1001)
	fmt.Printf("s3  len %d cap %d \n", len(s3), cap(s3))
}
