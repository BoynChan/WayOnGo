package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	//defer用于延迟处理,在defer归属的函数返回之前,defer声明的方法将会被一一执行
	//执行的顺序是先入后出,即一个栈
	fmt.Println("Start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("End")
	//defer语句可以在语句退出的时候执行,所以可以做一些释放资源的操作

	fmt.Println(fileSize("F:\\Code\\Go\\LearningGo\\src\\main\\Function\\Closure.go"))
}

func syncMapTest() {
	var (
		valueMap      = make(map[string]int)
		valueMapGuard sync.Mutex
	)

	valueMapGuard.Lock()
	//确保在方法返回的时候可以解锁
	defer valueMapGuard.Unlock()

	//对map的操作
	valueMap["someKey"] = 5
}

//根据文件名查询文件大小
func fileSize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return 0
	}
	return info.Size()
}
