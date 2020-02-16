package main

import (
	"fmt"
	"sync"
	"time"
)

/**
除了0405的同步锁之外,我们还可以使用Go语言中的一大特色
Channel
使用它在协程之间传输数据

这个程序使用无缓冲的通道模拟接力赛的情景
author:Boyn
date:2020/2/16
*/

var wg sync.WaitGroup

func main() {
	//创建一个无缓冲通道
	baton := make(chan int)

	wg.Add(1)
	go Runner(baton)
	//开始比赛
	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d Ready.\n", newRunner)
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("Runner %d Finished.\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}
