package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/**
创建一个程序用于生成指定目录的硬盘使用情况报告
即占用的大小
author:Boyn
date:2020/2/15
*/

func WalkDir(initDir []string) {
	fileSizes := make(chan int64)
	go func() {
		for _, root := range initDir {
			doWalkDir(root, fileSizes)
		}
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
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			doWalkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

/*
dirents返回目录下的文件
*/
func dirents(dir string) []os.FileInfo {
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
