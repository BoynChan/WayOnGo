package gee

import (
	"log"
	"net/http"
	"strings"
)

// Author:Boyn
// Date:2020/3/3

type HandlerFunc func(c *Context)

type RouteGroup struct {
	prefix      string        // 当前分组的前缀
	middlewares []HandlerFunc // 中间件集合
	parent      *RouteGroup   // 嵌套
	engine      *Engine       // 所有的RouteGroup都可以通过engine来调用资源
}

type Engine struct {
	*RouteGroup // engine作为最顶层的分组
	router      *route
	groups      []*RouteGroup // 存储所有的分组
}

func New() *Engine {
	engine := &Engine{router: newRoute()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (group *RouteGroup) Group(prefix string) *RouteGroup {
	engine := group.engine
	// 如果前缀没有 / ,则为其加上
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	// 如果后缀有 / ,则将其去除
	if strings.HasSuffix(prefix, "/") {
		prefix = prefix[0 : len(prefix)-1]
	}
	newGroup := &RouteGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouteGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouteGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Add Route %s - %s", method, comp)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouteGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouteGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	engine.router.handler(c)
}
