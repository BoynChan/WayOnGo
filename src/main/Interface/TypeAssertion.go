package main

import "fmt"

/*
类型断言,可以使用switch来判断类型
*/
func typeAssertion(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float32, float64:
		fmt.Println("float")
	}
}
func main() {
	typeAssertion(10)
	typeAssertion("hi")
	typeAssertion(5.12)
}
