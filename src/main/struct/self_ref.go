package main

func main() {

}

type Node struct {
	// nested 无法计算大小
	//left Node
	//right Node

	// 指针大小一致
	left  *Node
	right *Node
}
