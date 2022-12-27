package main

import (
	"fmt"
	"net/http"
	"strings"
)

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

// HandlerBasedOnTree ////////////////////////////////////////////////////////////////////
type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	// 如果这是叶子节点
	// 那么匹配上之后可以调用该方法
	Handler handlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	fmt.Printf("findRouter.... \n")
	// 去除头尾可能有的/，然后按照/切割成段
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		matchChild, found := h.findMatchChild(cur, p)
		// 从子节点里面查找一个匹配到当前path 的节点
		//matchChild, found := cur.findMatchChildV2(p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	if cur.Handler == nil {
		// 注册了 /user/profiler
		// 然问了 /user
	}
	return cur.Handler, true
}

func (h *HandlerBasedOnTree) Route(method string, pattern string,
	handleFunc handlerFunc) {
	fmt.Printf("Route.... \n")
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root
	for index, path := range paths {
		// 通过根结点找
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		} else {
			fmt.Printf("goto createSubTree.... \n")
			// 通过根结点找
			h.createSubTree(cur, paths[index:], handleFunc)
			return
		}
	}
}

// demo:
// user/profiler/setting
// user/home
// user/friends
// order/detail

// paths 可以是 friends/tom/address
func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handleFn handlerFunc) {
	fmt.Printf("createSubTree.... \n")
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		// user.children = [profile, home, friends]
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.Handler = handleFn
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 8),
	}
}

func (h *HandlerBasedOnTree) findMatchChild(root *node,
	path string) (*node, bool) {
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}

	return nil, false
}

var _ Handler = &HandlerBasedOnTree{}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBasedOnTree{
		root: &node{},
	}
}
