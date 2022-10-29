package main

import "fmt"

func main() {
	// 结构体，所以方法调用的时候它的数据不会变
	u := User{
		Name: "Tom",
		Age:  18,
	}

	u.ChangeName("TOM changed")
	u.ChangeAge(20)
	fmt.Printf("%v \n", u)
	// {Tom 20}

	// 指针，所以内部的数据是可以改变的
	up := &User{
		Name: "Jerry",
		Age:  21,
	}

	up.ChangeName("Jerry changed")
	up.ChangeAge(23)
	fmt.Printf("%v \n", up)
	// &{Jerry 23}
}

type User struct {
	Name string
	Age  int
}

// 结构体接收器
func (u User) ChangeName(newName string) {
	//http.Handle()
	u.Name = newName
}

// 指针接收器 recommend 1
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

// recommend 2
type Handle func()

func (h Handle) Hello() {

}
