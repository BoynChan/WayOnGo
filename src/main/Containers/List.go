package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list的底层是双链表
	l := list.New()
	//在队尾插入
	l.PushBack(1)
	//在队头插入
	l.PushFront(5)

	//用于保存句柄,此时element是新的队尾,即10
	element := l.PushBack(10)

	//在element之前插入一个值,可以看得出,9会在1之后,10之前
	l.InsertBefore(9, element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
