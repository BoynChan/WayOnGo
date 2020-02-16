package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/**
根据03程序中竞态条件,我们可以使用atomic和sync中的函数
来对共享资源加锁
在这个程序中就使用了atomic对counter进行原子相加
author:Boyn
date:2020/2/16
*/

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter2(1)
	go incCounter2(2)
	wg.Wait()
	fmt.Println(counter)
}

func incCounter2(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
