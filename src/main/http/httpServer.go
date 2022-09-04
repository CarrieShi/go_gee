package main

import "net/http"

type Server interface {
	// Route 设定路由，命中路由会执行 handlerFunc 的代码
	Route(pattern string, handleFunc http.HandlerFunc)

	// Start 启动服务
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

type Header map[string][]string
