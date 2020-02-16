package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
分别打印大写和小写的英文字母
对该程序来说,由于有两个逻辑处理器
所以Goroutines的调度也会变为并行的方式
每一个Goroutines运行在自己的核中
author:Boyn
date:2020/2/15
*/
func main() {
	//表示分配2个逻辑处理器
	//这会造成,字母表的大写和小写是乱序打印的
	runtime.GOMAXPROCS(2)

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
