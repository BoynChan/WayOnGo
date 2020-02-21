package main

import (
	"fmt"
	"html/template"
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

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("F:\\Code\\Go\\LearningGo\\src\\web\\goweb\\page\\login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Printf("username:%s\n", r.Form["username"])
		fmt.Printf("password:%s\n", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8009", nil)
	if err != nil {
		log.Fatal(err)
	}
}
