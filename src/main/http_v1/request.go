package main

import (
	"net/http"
)

func main() {
	server := NewHttpServer("test-server", MetricsFilterBuilder)

	server.Route(http.MethodPost, "/user/signUp", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
