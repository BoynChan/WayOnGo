package main

import (
	"flag"
	"fmt"
	"time"
)

//接收命令行参数,指定秒数,休眠特定的事件
var period = flag.Duration("period", 1*time.Second, "sleep period")

// go run ./Sleep.go -period=2s
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", period)
	time.Sleep(*period)
	fmt.Println()
}
