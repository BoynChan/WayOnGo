package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

/**
创建一个程序用于生成指定目录的硬盘使用情况报告
即占用的大小
author:Boyn
date:2020/2/15
*/

func WalkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			WalkDir(subdir, fileSizes)
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
