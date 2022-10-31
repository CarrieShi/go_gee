package main

import "net/http"

type HandlerBasedOnMap struct {
	// key = method + url
	// val = handleFunc 类型 func(ctx *Context)
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		// 找到了注册的路径
		handler(NewContext(writer, request))
	} else {
		// 找不到 404
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Method Not allowed"))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}
