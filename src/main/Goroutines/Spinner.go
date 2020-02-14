package main

import (
	"fmt"
	"time"
)

/**
计算斐波那契数列的同时,让用户看到一个可见的标识来表明程序运行正常
author:Boyn
date:2020/2/14
*/
func main() {
	go spinner(100 * time.Millisecond)
	// 计算第45个斐波那契数
	const n = 45
	fibN := fib(n)
	fmt.Println("Fibonacci:", fibN)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-/\\\\|" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
