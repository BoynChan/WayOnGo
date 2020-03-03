package gee

import (
	"log"
	"net/http"
	"strings"
)

// Author:Boyn
// Date:2020/3/3
// 将Handler函数作为一个类型,后面方便引用

type route struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRoute() *route {
	return &route{handlers: make(map[string]HandlerFunc), roots: make(map[string]*node)}
}

// 解析pattern
// 将传入的路径用 / 进行切分
// 一个一个加入到表示路径的数组中
// 当遇到 * 时停止并返回
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 将外部注册的路径放入route中
// 首先进行路径解析
// 然后根据对应方法,找到route中前缀树的根节点
// 将该路径模式放入请求方法对应的前缀树中
func (r *route) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Add Route %4s - %s", method, pattern)
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// 返回路由表中匹配的路径对应的node节点以及动态配置参数
// 首先对传入的路径进行解析,并获取对应方法的前缀树根节点
// 以此根节点为起点进行路径匹配的搜索,并返回匹配的节点n
// 当节点不为空的时候,解析这个n对应的路径(我们传入的路径是不带动态参数的,而n对应的路径在传入时可以有动态参数)
// 将解析出来的路径与我们传入的路径逐个对比,并将动态参数放入params中
// 比如我们传入/search/go/sth 而动态路径是 /search/:lang/sth 那么结果就会返回当前节点与{lang:"go"}的map
// 而如果传入/static/css/gee.css 而动态路径是 /static/*filepath 那么结果就会返回当前节点和{filepath:"css/gee.css"}的map
func (r *route) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *route) handler(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s", c.Path)
	}
}
