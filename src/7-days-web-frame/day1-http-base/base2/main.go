package main

import (
	"fmt"
	"log"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3

// 在 base1 中,我们使用了http包自带的http接口
// 但是在 ListenAndServe 函数中,我们是可以传入自己的一个Handler来实现更多定制化的功能的

//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}

// Engine是所有请求的handler
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		_, _ = fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q], = %q\n", k, v)
		}
	default:
		_, _ = fmt.Fprintf(w, "404 NOT FOUNT Path=%q\n", req.URL.Path)
	}
}

func main() {
	engine := &Engine{}
	log.Fatal(http.ListenAndServe(":8899", engine))
}
