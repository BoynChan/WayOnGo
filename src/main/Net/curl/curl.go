package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
使用io.Reader与io.Writer接口
编写一个简单的curl程序
author:Boyn
date:2020/2/17
*/

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./curl <url>")
		os.Exit(-1)
	}
}

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
