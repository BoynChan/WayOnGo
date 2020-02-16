package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
除了0405的同步锁之外,我们还可以使用Go语言中的一大特色
Channel
使用它在协程之间传输数据

这个程序模拟一个网球的击打游戏
author:Boyn
date:2020/2/16
*/
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建了一个无缓冲的通道
	count := make(chan int)

	// 计数+2
	wg.Add(2)

	go player("Bob", count)
	go player("Jack", count)
	//表示发球的动作
	count <- 1
	wg.Wait()
	fmt.Println("Game end")
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
