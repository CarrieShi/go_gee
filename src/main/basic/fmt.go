package main

import "fmt"

func replaceHolder() {
	u := &user{
		name: "test",
		age:  18,
	}

	fmt.Printf("v => %v \n", u)
	fmt.Printf("+v => %+v \n", u)
	fmt.Printf("#v => %#v \n", u)
	fmt.Printf("T => %T \n", u)
}

type user struct {
	name string
	age  int
}

func main() {
	replaceHolder()
}
