package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("cat", "Tom")
	m.Store("mouse", "Jerry")

	// 这里需要重新读取
	// 1.类型断言 t, ok := x.(T)
	// T 可以是结构体 或者 指针
	// 类似 instanceOf + 强制类型转换
	// x = nil 返回false
	// 编译器不会检查
	// 2.区别于类型转化 y := T(x)
	// 类似于强制转换
	// 编译器会进行类型检查，不能转换的会编译错误
	val, ok := m.Load("cat")
	if ok {
		fmt.Println(len(val.(string)))
	}
}
