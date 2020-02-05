package main

import "fmt"

func main() {
	var car int = 1
	var ba string = "baba"
	fmt.Printf("%p,%p\n", &car, &ba) //显示两个变量的地址

	house := "A house refers to a string"
	//对字符串取地址,ptr的类型为*string
	ptr := &house
	fmt.Printf("ptr type : %T\n", ptr)

	fmt.Printf("ptr address : %p\n", ptr)

	//对指针进行取值操作,value的值与house的值相同
	value := *ptr

	fmt.Printf("value type : %T\n", value)

	// 在这里 house和value在内存中的地址是不一样的,这里体现当我们用一个变量来对指针取值的时候
	// 其实是在申请一个新的变量,将其赋值为指针所对应内存地址的值
	fmt.Printf("value & house address : %p,%p\n", &value, &house)

	//用指针进行数值交换
	a, b := 5, 6
	fmt.Printf("before swap a,b %d,%d \t %p,%p\n", a, b, &a, &b)
	swap(&a, &b)
	//可以看到我们的交换并不是交换地址,而是取到地址后交换他们的值
	fmt.Printf("after swap a,b %d,%d \t %p,%p\n", a, b, &a, &b)

	//还可以用new操作符来创建指针
	str := new(string)
	*str = "Go语言"
	fmt.Println(*str)
}

func swap(a, b *int) {
	//取a指针的值
	t := *a
	*a = *b
	*b = t
}
