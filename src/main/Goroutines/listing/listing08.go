package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
展示如何使用有缓冲的通道和固定数目的goroutine 来处理数据
对于通道来说,如果有多个goroutines尝试获取通道中的内容,那么获取的顺序通常是随机的
所有每次运行的结果都会不一样
author:Boyn
date:2020/2/16
*/

const (
	numberGoroutines = 4  // 使用的goroutines的数量
	taskLoad         = 10 // 要处理的工作的数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// 创建了一个缓冲大小的taskLoad的通道进行处理
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)

	// 启动Goroutines来处理工作
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完成后关闭通道
	// 以便所有 goroutines 退出
	// 通道关闭后仍然可以从中拿取数据,但是不能从中写入数据
	close(tasks)

	wg.Wait()

}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			//意味着通道被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker %d start work for Task:%s\n", worker, task)

		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		fmt.Printf("Worker %d finished work for Task:%s\n", worker, task)
	}
}
