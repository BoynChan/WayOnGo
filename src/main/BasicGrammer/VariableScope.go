package main

/**
变量作用域
*/
import "fmt"

//定义一个int值为10的全局变量
var c int
var d float32 = 0.18

func main() {
	a := 20
	b := 30
	c = a + b
	//在变量重名时,局部变量优先于全局变量
	d := 5
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
}

/**
a,b是函数中形式参数,在函数未调用的时候不会占用实际存储空间
在调用时被赋值,调用完成后销毁
*/
func sum(a, b int) int {
	return a + b
}
