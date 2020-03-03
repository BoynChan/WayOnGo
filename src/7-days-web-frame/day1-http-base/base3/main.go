package main

import (
	"7-days-web-frame/day1-http-base/base3/gee"
	"fmt"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3

func main() {
	engine := gee.New()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	})

	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			_, _ = fmt.Fprintf(w, "Header[%q], = %q\n", k, v)
		}
	})

	engine.Run(":8899")

}
