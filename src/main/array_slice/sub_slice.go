package main

import "fmt"

func main() {
	s1 := make([]int, 0, 10)
	fmt.Printf("s1 %v len %d cap %d \n", s1, len(s1), cap(s1))

	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("s1 %v len %d cap %d \n", s1, len(s1), cap(s1))

	subS1 := s1[2:3] // index > 2 && index <= 3
	fmt.Printf("subS1 %v len %d cap %d \n", subS1, len(subS1), cap(subS1))
	// subS1 [3] len 1 cap 8

	subS2 := s1[:4]
	fmt.Printf("subS2 %v len %d cap %d \n", subS2, len(subS2), cap(subS2))

	subS3 := s1[2:]
	fmt.Printf("subS3 %v len %d cap %d \n", subS3, len(subS3), cap(subS3))

	// 子切片和原切片共享数据？？？
	subS1 = append(subS1, 11)
	subS1 = append(subS1, 12, 13, 14, 15, 16, 17)
	fmt.Printf("subS1 %v len %d cap %d \n", subS1, len(subS1), cap(subS1))
	// subS1 [3 11 12 13 14 15 16 17] len 8 cap 8

	subS1 = append(subS1, 18, 19, 20, 21, 22)
	fmt.Printf("subS1 %v len %d cap %d \n", subS1, len(subS1), cap(subS1))
	// subS1 [3 11 12 13 14 15 16 17 18 19 20 21 22] len 13 cap 16

	fmt.Printf("s1 %v len %d cap %d \n", s1, len(s1), cap(s1))
	// s1 [1 2 3 11 12 13 14 15 16 17] len 10 cap 10
}
