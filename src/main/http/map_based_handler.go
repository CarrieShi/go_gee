package main

import "net/http"

// Routable httpServer 也会路由，故抽象出 Routable 接口
type Routable interface {
	RouteBasedOnMethod(
		method string,
		pattern string,
		handleFunc func(ctx *Context))
}

// Handler 使用组合，解耦，RouteBasedOnMethod 接口
type Handler interface {
	//http.Handler
	ServeHTTP(c *Context)
	Routable
	//RouteBasedOnMethod(
	//	method string,
	//	pattern string,
	//	handleFunc func(ctx *Context))
}

type HandlerBasedOnMap struct {
	// key = method + url
	// val = handleFunc 类型 func(ctx *Context)
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

// RouteBasedOnMethod 解强耦合 移动至此 实现
func (h *HandlerBasedOnMap) RouteBasedOnMethod(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		// 找到了注册的路径
		handler(c)
	} else {
		// 找不到 404
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Method Not allowed"))
	}
}

//func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter,
//		request *http.Request) {
//	key := h.key(request.Method, request.URL.Path)
//	if handler, ok := h.handlers[key]; ok {
//		// 找到了注册的路径
//		handler(NewContext(writer, request))
//	} else {
//		// 找不到 404
//		writer.WriteHeader(http.StatusNotFound)
//		writer.Write([]byte("Method Not allowed"))
//	}
//}

// 确保 HandlerBasedOnMap 一定实现了 Handler ???
// 如果 Handler 新增了方法，HandlerBasedOnMap 没有实现的情况下，ide 会报错
// 看上去像声明 初始化
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBaseOnMap() Handler {
	return &HandlerBasedOnMap{
		//handlers: make(map[string]func(c *Context), 128), // 预估容量，不够可自动扩容
		handlers: make(map[string]func(c *Context)), // 预估容量，不够可自动扩容
	}
}
