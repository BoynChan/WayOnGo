package main

import (
	"fmt"
	"sort"
)

/*
介绍原生的排序接口
*/
type String []string

//下面的Len() Less() Swap()方法都是通过实现sort.Interface接口从而实现了排序需要的要素
//从而可以调用标准的排序函数进行排序
func (s String) Len() int {
	return len(s)
}

func (s String) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s String) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	names := String{
		"Jose",
		"Bob",
		"Channel",
		"Kit",
		"John",
		"Rose",
		"Jack",
	}
	sort.Sort(names)
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
