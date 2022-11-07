package main

import "fmt"

func main() {
	// defer 延迟执行
	defer func() {
		// recover 内置方法
		if data := recover(); data != nil {
			fmt.Printf("现场发生了什么？！ %v \n", data)
		}
		fmt.Printf("信号恢复连接，继续 \n")
	}()

	panic("Boom Boom Boom")
	fmt.Println("爆炸后的现场，信号中断")
}
