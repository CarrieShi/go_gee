package main

import "unicode/utf8"

func main() {
	println("hello go")
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))
	println(`
{"json":"test"}
`)
}
