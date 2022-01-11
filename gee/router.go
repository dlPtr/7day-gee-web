package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)
type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node // 基于请求方式存储，允许不同请求方式执行不同的回调函数
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, part := range vs {
		if part == "" {
			continue
		}

		parts = append(parts, part)

		// 碰到了通配符 * ，则后续路径直接忽略
		if part[0] == '*' {
			break
		}
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	// 插入路由，映射：请求路径(method+_+pattern)->处理回调
	key := method + "-" + pattern
	r.handlers[key] = handler

	// 记录已经存在的请求路径 pattern 到 trie 树
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
}

// @brief: 根据请求路径，在tries树中匹配对应的pattern
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	parts := parsePattern(n.pattern)
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
			break
		}
	}
	return n, params
}

func (r *router) handle(c *Context) {
	if n, params := r.getRoute(c.Method, c.Path); n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "<h1>404 NOT FOUND: %s</h1>", c.Path)
	}
}
