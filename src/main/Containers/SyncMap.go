package main

import (
	"fmt"
	"sync"
)

func main() {
	/*//创建一个map,先来看看并发场景下会有什么问题
	m := make(map[int]int)

	//开启一段并发的写入代码,不断对map进行写入
	go func() {
		for{
			m[1] =1
		}
	}()

	//不断读取
	go func() {
		for{
			_ = m[1]
		}
	}()*/

	/*for{

	}*/
	//结果是会报错
	// fatal error: concurrent map read and map write

	/*
		所以,我们需要使用在并发情况下
		线程安全的Map -- sync.Map
		他有以下特性,无需初始化,直接声明
		不能以map的方式进行取值或者设置等操作,要使用他自己的方法
		使用Range配合一个回调函数进行遍历操作
	*/

	var syncMap sync.Map
	syncMap.Store("One", 1)
	syncMap.Store("Two", 2)
	syncMap.Store("Three", 3)
	syncMap.Store("Ten", 10)

	fmt.Println(syncMap.Load("One"))
	syncMap.Delete("Ten")
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
