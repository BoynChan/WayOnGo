package main

import "fmt"

/**
复数
*/

func main() {
	var a complex128 = complex(1, 2) //1+2i
	b := complex(3, 4)               //2+1i
	fmt.Println(a * b)               //(-5+10i)
}
