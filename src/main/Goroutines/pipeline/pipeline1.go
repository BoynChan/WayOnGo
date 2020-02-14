package pipeline

import (
	"fmt"
	"time"
)

/**
将多个channel串联在一起,一个channel的输出作为下一个的输入
author:Boyn
date:2020/2/14
*/

/*
第一个流水线是一个计数器+平方器+打印器
*/
func pipeline1() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
			//避免速度过快,进行睡眠
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // 如果ok为false,即说明不能从通道中接收到值,通道已经关闭了,所以退出循环
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}
