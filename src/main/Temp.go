package main

import "fmt"

/**
a,b互换
可以直接赋值来转换
*/
func main() {
	a := 100
	b := 200
	b, a = a, b
	fmt.Printf("%d %d\n", a, b)
}
