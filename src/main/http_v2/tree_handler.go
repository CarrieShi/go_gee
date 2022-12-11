package main

import (
	"net/http"
	"strings"
)

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
	handleFunc handlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root
	for index, path := range paths {
		matchChild, ok := cur.findMatchChild(path)
		if ok {
			cur = matchChild
		} else {
			cur.createSubTree(paths[index:], handleFunc)
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
func (n *node) createSubTree(paths []string, handleFn handlerFunc) {
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

func (n *node) findMatchChild(path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range n.children {
		if child.path == path &&
			child.path != "*" {
			return child, true
		}
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
			// todo:
			// ErrorInvalidRouterPattern
		}
		if pattern[pos-1] != '/' {
			// todo:
		}
	}
}

var _ Handler = &HandlerBasedOnTree{}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBasedOnTree{
		root: &node{
			path: "", // todo:
		},
	}
}
