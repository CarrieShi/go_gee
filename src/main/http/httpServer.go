package main

import "net/http"

type Server interface {
	// Route 设定路由，命中路由会执行 handlerFunc 的代码
	Route(pattern string, handleFunc http.HandlerFunc)

	// Start 启动服务
	Start(address string) error
}

// sdkHttpServer 基于 http 库实现
type sdkHttpServer struct {
	Name string
}

// Route 注册路由
func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (s *sdkHttpServer) Start(address string) error {
	//TODO implement me
	panic("implement me")
}
