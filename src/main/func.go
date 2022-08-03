package main

func returnTest(code int, errMsg string) (int, string) {
	return code, errMsg
}

func returnMany() (code, bcode int, msg string) {
	return 404, -1, "not found 404-1"
}

func returnWithName() (age int, name string) {
	age = 10
	name = "test-name"
	//other = "male"
	return
}

func main() {
	code, msg := returnTest(404, "not found")
	println(code, msg)

	code, _, msg = returnMany()
	println(code, msg)

	age, name := returnWithName()
	println(age, name)
}
