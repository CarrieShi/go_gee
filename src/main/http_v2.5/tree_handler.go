package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Routable interface {
	Route(
		method string,
		pattern string,
		handlerFunc handlerFunc) error
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
		// 从子节点里面查找一个匹配到当前path 的节点
		matchChild, found := cur.findMatchChild(p)
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
	handlerFunc handlerFunc) error {
	fmt.Printf("Route.... \n")

	// 校验模式
	err := h.validatePattern(pattern)
	if err != nil {
		return err
	}

	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root
	for index, path := range paths {
		// 通过节点找
		matchChild, ok := cur.findMatchChild(path)
		if ok {
			cur = matchChild
		} else {
			fmt.Printf("goto createSubTree.... \n")
			// 通过节点找
			cur.createSubTree(paths[index:], handlerFunc)
			return nil
		}
	}

	// 离开了循环，加入短路径
	// 如 先加入 /order/detail
	// 再加入 /order 就会走这里
	cur.Handler = handlerFunc
	return nil
}

// demo:
// user/profiler/setting
// user/home
// user/friends
// order/detail

// paths 可以是 friends/tom/address
func (n *node) createSubTree(paths []string, handleFn handlerFunc) {
	fmt.Printf("createSubTree.... \n")
	cur := n
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

// 注册了/order/*
// 又注册了/order/*/confirm ==》设计不允许，收益小，风险大，难维护
// 请求 /order/123/cancel 可以命中 /order/* 吗？
// ==》设计不可以，如果可以会很麻烦，到了cancel 层还需要回溯
func (n *node) findMatchChild(path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range n.children {
		// 优先挑选详细的
		// != * 是防止用户乱输入
		if child.path == path &&
			child.path != "*" {
			return child, true
		}
		// 命中了通配符的，后面看下有没有更详细的
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}

/*func (n *node) findMatchChild(path string) (*node, bool) {
	for _, child := range n.children {
		if child.path == path {
			return child, true
		}
	}

	return nil, false
}*/

func (h *HandlerBasedOnTree) validatePattern(pattern string) error {
	// 校验 * 是否存在，如果存在，必须在最后一个，并且它前面必须是 /
	// 即 只接受 /* 的存在，abc*、**、*/aaa/bbb 是非法的

	pos := strings.Index(pattern, "*")
	if pos > 0 {
		if pos != len(pattern)-1 {
			return errors.New("ErrorInvalidRouterPattern")
		}
		if pattern[pos-1] != '/' {
			return errors.New("ErrorInvalidRouterPattern")
		}
	}
	return nil
}

var _ Handler = &HandlerBasedOnTree{}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBasedOnTree{
		root: &node{},
	}
}
