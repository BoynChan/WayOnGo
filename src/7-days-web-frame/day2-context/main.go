package main

import (
	"7-days-web-frame/day2-context/gee"
)

// Author:Boyn
// Date:2020/3/3

func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>Hello Gee</h1>")
	})
	engine.GET("/hello", func(c *gee.Context) {
		name := c.Query("name")
		c.String(200, "Hello %s", name)
	})
	engine.Run(":8899")
}
