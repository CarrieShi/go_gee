package main

import (
	"net/http"
)

type Server interface {
	Routable
	start(address string) error
}

// sdkHttpServer 基于 http 库实现
type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnTree()
	var root Filter = handler.ServeHTTP

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}
