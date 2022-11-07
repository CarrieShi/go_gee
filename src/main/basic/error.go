package main

import (
	"errors"
	"fmt"
)

func main() {
	//var err error = &MyError{}
	//println(err.Error())

	//err2 := errors.New("new a error")
	//println(err2.Error())

	ErrorPkg()
}

type MyError struct {
}

func (m *MyError) Error() string {
	panic("this is my error")
}

func ErrorPkg() {
	err := &MyError{}
	// %w 占位符，返回一个新的错误
	// wrappedErr 是一个新类型， fmt.wrapError
	// fmt.PrintXXX to console
	// fmt.SprintXXX return string
	// fmt.ErrorXXX error
	wrappedErr := fmt.Errorf("this is an wrapped error %w", err)

	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("unwrapped")
	}

	if errors.Is(wrappedErr, err) {
		// is 会逐层解除包装，判断是不是该错误
		fmt.Println("wrapped is err")
	}

	copyErr := &MyError{}
	// 将 wrappedErr 转为 MyError
	// 使用两次取地址符
	if errors.As(wrappedErr, &copyErr) {
		fmt.Println("convert error")
	}
}
