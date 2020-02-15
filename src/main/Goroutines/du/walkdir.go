package du

import (
	"fmt"
	"io/ioutil"
	"main/Goroutines/semaphore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/**
创建一个程序用于生成指定目录的硬盘使用情况报告
即占用的大小
author:Boyn
date:2020/2/15
*/
var waitGroup sync.WaitGroup
var s = semaphore.NewSemaphore(20)

func WalkDir(initDir []string) {

	fileSizes := make(chan int64)

	for _, root := range initDir {
		//增加一个计数的值
		waitGroup.Add(1)
		go doWalkDir(root, fileSizes)
	}
	go func() {
		//当n的值为0时,就会结束阻塞
		waitGroup.Wait()
		close(fileSizes)
	}()

	// 为了显示进度,所以每500ms显示一次收集进度
	tick := time.Tick(500 * time.Millisecond)

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
	fmt.Println()
}

func doWalkDir(dir string, fileSizes chan<- int64) {
	//在函数结束的时候,会将等待组的值减一
	defer waitGroup.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			//如果是目录,会进入递归函数,所以将等待组的值加一
			waitGroup.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			doWalkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

/*
dirents返回目录下的文件
在并发过程中,为了防止并发数量过多
在dirent函数中使用一个信号量作为并发量的控制
*/
func dirents(dir string) []os.FileInfo {
	//信号量
	s.Acquire()
	defer s.Release()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files : %.1f MB Uses\n", nfiles, float64(nbytes)/1e6)
}
