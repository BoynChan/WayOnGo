package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
这个程序使用了mutex加锁来划定一个临界区
以此来消除竞态条件
author:Boyn
date:2020/2/16
*/

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println(counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
