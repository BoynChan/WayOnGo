package main

import "fmt"

/**
匿名变量
并且可以用_来接收函数中返回值不需要的那一部分
匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
*/
func main() {
	a, _ := GetDate()
	fmt.Printf("%d ", a)
	_, b := GetDate()
	fmt.Printf("%d ", b)
}

func GetDate() (int, int) {
	return 100, 200
}
