package _select

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

/**
增加了按回车键结束发射过程的功能
author:Boyn
date:2020/2/15
*/

func countdown2() {
	abort := make(chan int)
	go func() {
		//创建一个Scanner,当没有输入时,会一直阻塞,当有任意输入的时候,就结束阻塞
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println("stop")
		abort <- 1
	}()

	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second) //Tick函数返回一个阻塞channel,周期性地发送一个时间
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		// 什么也不干
		case <-abort:
			fmt.Println("Abort")
			//当接收到abort信息的时候,说明回车被按下,则发射过程终止
			return
		}
	}
	launch()
}
