package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
这个程序包含竞态条件
在这个程序中,value++被调用了4次,理应counter的最终值也是4
但是counter最终结果是2
因为每个goroutines都有自己的空间,在他们的空间中
value的值是2,那么写回counter的时候,也是2
可以使用go build -race  .\src\main\Goroutines\listing\listing03.go
并运行程序来检测竞态条件

author:Boyn
date:2020/2/16
*/

var (
	counter int64
	wg      sync.WaitGroup
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
		value := counter
		//运行该函数表示当前goroutine从线程中退出,放回到调度队列中
		runtime.Gosched()

		value++

		counter = value
	}
}
