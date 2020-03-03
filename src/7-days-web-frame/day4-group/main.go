package main

import (
	"7-days-web-frame/day4-group/gee"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee!</h1>")
	})
	v1 := r.Group("/v1")
	v1.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>V1 :: Hello Gee!</h1>")
	})
	v2 := r.Group("v2")
	v2.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>V2 :: Hello Gee!</h1>")
	})

	r.Run(":8899")
}
