package main

import (
	"net/http"
)

type Server interface {
	// RouteRW 设定路由，命中路由会执行 handlerFunc 的代码
	RouteRW(pattern string, handleFunc http.HandlerFunc)
	Route(pattern string, handleFunc func(ctx *Context))

	// Start 启动服务
	Start(address string) error
}

// sdkHttpServer 基于 http 库实现
type sdkHttpServer struct {
	Name string
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

////////////////////// Route 注册路由 对比开始 //////////////////////

func (s *sdkHttpServer) RouteRW(pattern string,
	handleFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handleFunc)
}

// Route 使用 Context 代替 type HandlerFunc func(ResponseWriter, *Request)
func (s *sdkHttpServer) Route(pattern string,
	handleFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter,
		request *http.Request) {
		ctx := NewContext(writer, request)
		handleFunc(ctx)
	})
}

////////////////////// Route 注册路由 对比结束 //////////////////////
