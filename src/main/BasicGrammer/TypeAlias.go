package main

import "fmt"

//这个是定义内建类型
type byte1 uint8

// 这个是类型别名
type byte2 = uint8

func main() {
	//对于两种类型的定义来说
	//定义自建类型是创造一个新的类型,但仍然具有原类型的特性
	//而类型别名则是像C中的宏替换
	var a byte1 = 10
	var b byte2 = 10
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
