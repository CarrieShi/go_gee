package main

import (
	"net/http"
)

type Server interface {
	// RouteRW 设定路由，命中路由会执行 handlerFunc 的代码
	//RouteRW(pattern string, handleFunc http.HandlerFunc)
	//Route(pattern string, handleFunc func(ctx *Context))

	// Start 启动服务
	//Start(address string) error

	// RESTful
	RouteBasedOnMethod(method string, pattern string, handleFunc func(ctx *Context))
	StartBasedOnMethod(address string) error
}

// sdkHttpServer 基于 http 库实现
type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedOnMap // 声明 // 强耦合
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
		// 初始化
		handler: &HandlerBasedOnMap{
			handlers: map[string]func(ctx *Context){},
		},
	}
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func (s *sdkHttpServer) StartBasedOnMethod(address string) error {
	http.Handle("/", s.handler)
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

func (s *sdkHttpServer) RouteBasedOnMethod(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	// 仅对 HandlerBasedOnMap handler 的赋值
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handleFunc

	// panic: assignment to entry in nil map
	// 解决：NewHttpServer 内需初始化 sdkHttpServer.handler
	// 否则 它的值是nil, 不指向任何内存地址，即报上诉的 panic

	// todo: 检测重复注册问题

	// handler := &HandlerBasedOnMap{}
	// 给 handler 赋值...

	// HandlerBasedOnMap 只需要注册一遍
	// 在 sdkHttpServer 声明
	// 在 StartBasedOnMethod 调用
	// http.Handle("/", handler)
}

////////////////////// Route 注册路由 对比结束 //////////////////////
