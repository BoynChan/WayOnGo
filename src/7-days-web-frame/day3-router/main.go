package main

import (
	"7-days-web-frame/day3-router/gee"
	"net/http"
)

// Author:Boyn
// Date:2020/3/3

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee!</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(200, "Hello! %s", c.Query("name"))
	})
	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(200, "Hello! %s", c.Param("name"))
	})
	r.GET("/static/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":8899")
}
