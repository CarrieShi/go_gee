package main

import "fmt"

func main() {
	// 创建预估容量2的 map
	m := make(map[string]string, 2)
	// 创建没有预估容量的 map
	m1 := make(map[string]string)
	// 直接初始化
	m2 := map[string]string{
		"name": "Tom",
	}

	// 赋值
	m["name"] = "Jerry"
	m1["name"] = "Jim"
	m2["gender"] = "male"

	// 取值
	val := m["name"]
	println(val)

	// 取值，返回两个值，是否存在该key
	val, ok := m1["invalid_key"]
	if !ok {
		println("key not found")
	}

	// map key 顺序不定
	for key, val := range m {
		fmt.Printf("#{%s} ==> #{%s} \n", key, val)
	}

}
