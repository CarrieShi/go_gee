package main

import "net/http"

type Routable interface {
	Route(
		method string,
		pattern string,
		handleFunc handlerFunc)
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

// HandlerBasedOnMap ////////////////////////////////////////////////////////////////////
type HandlerBasedOnMap struct {
	// key = method + url
	// val = handleFunc 类型 func(ctx *Context)
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func (h *HandlerBasedOnMap) Route(
	method string,
	pattern string,
	handleFunc handlerFunc) {
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

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBaseOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(c *Context)), // 预估容量，不够可自动扩容
	}
}
