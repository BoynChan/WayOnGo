package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
分别打印大写和小写的英文字母
对该程序来说,由于只有一个逻辑处理器,所以执行的顺序是可以预知的
首先我们知道main函数也是一个Goroutines,所以一开始会先执行main协程
然后在协程中声明的go func语句会以栈作为存储结构,在main协程阻塞等待或被调度的时候
会执行后面声明的go func,再执行前面执行的go func
author:Boyn
date:2020/2/15
*/
func main() {
	//表示仅分配一个逻辑处理器
	//这会造成,就算有多个处理器,也不会发生并行处理的情况
	//便于我们进行交替打印
	runtime.GOMAXPROCS(1)

	// wg用于等待程序的完成
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()
		//显示小写字母表
		for char := 'a'; char <= 'z'; char++ {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()
		//显示大写字母表
		for char := 'A'; char <= 'Z'; char++ {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}()
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Done")
}
