package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3} //3个元素的数组
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	b := [...]int{1, 2, 3, 4, 5} //根据后面初始化的个数来确定数组大小...表示不确定
	for index, value := range b {
		fmt.Printf("index:%d value:%d\n", index, value)
	}

	//创建一个字符串数组
	course := [...]string{
		"database",
		"network",
		"Java",
		"Go"}

	for _, v := range course {
		fmt.Printf("course:%s\n", v)
	}

	//创建一个2X2的数组
	array := [2][2]int{{1, 2}, {3, 4}}
	for _, level := range array {
		for _, v := range level {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}

}
