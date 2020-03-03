package main

import (
	"fmt"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3

// 创建一个最基本的http服务,使用go自带的http接口

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8898", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
}
