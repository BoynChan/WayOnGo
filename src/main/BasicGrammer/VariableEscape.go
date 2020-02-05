package main

import "fmt"

/**
变量逃逸
*/

func dummy(b int) int {
	var c int
	c = b
	return c
}
func void() {}

/**
运行结果
go run -gcflags "-m -l" .\src\main\VariableEscape.go
# command-line-arguments
src\main\VariableEscape.go:18:13: main ... argument does not escape
src\main\VariableEscape.go:18:13: a escapes to heap
src\main\VariableEscape.go:18:21: dummy(0) escapes to heap
0 0

其中c变量是整形,其作用域由于dummy的返回逃逸出了dummy()函数
*/
func main() {
	var a int
	void()
	fmt.Println(a, dummy(0))
}
