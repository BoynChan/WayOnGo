package gee

import (
	"log"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3
// 将Handler函数作为一个类型,后面方便引用

type route struct {
	handlers map[string]HandlerFunc
}

func newRoute() *route {
	return &route{handlers: make(map[string]HandlerFunc)}
}

func (r *route) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Add Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *route) handler(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}
