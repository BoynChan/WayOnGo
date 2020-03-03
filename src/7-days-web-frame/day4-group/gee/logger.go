package gee

import (
	"log"
	"time"
)

// Author:Boyn
// Date:2020/3/3

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
