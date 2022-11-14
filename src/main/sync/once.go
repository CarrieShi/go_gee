package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

var once sync.Once

func PrintOnce() {
	once.Do(func() {
		fmt.Printf("人不能n次踏进同一条河流(n>=2) \n")
	})
}
