package gee

import (
	"fmt"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3

// 将Handler函数作为一个类型,后面方便引用
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// 实现了路由表,即路径与Handler函数的映射
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET方法被外部调用,添加一个HTTP类型的GET的请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUNT : %s\n", req.URL)
	}
}
