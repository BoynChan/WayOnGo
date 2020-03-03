package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter // 响应
	Req        *http.Request       // 请求
	Path       string              // 请求的路径
	Method     string              // 请求的方法
	Params     map[string]string   // 动态路径的解析参数
	StatusCode int                 // 响应状态码
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	_, _ = fmt.Fprintf(c.Writer, format, value...)
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}
