package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Author:Boyn
// Date:2020/2/21

// 使用go自带的http包创建一个简单的Web服务器

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数
	fmt.Println(r.Form) // 打印http请求的参数与数据
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Printf("%s:%s\n", k, strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":8009", nil)
	if err != nil {
		log.Fatal(err)
	}
}
