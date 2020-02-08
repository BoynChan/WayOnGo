package main

import "fmt"

/*
有如下的这么一个接口
这个接口需要实现Call方法
Call方法在调用的时候会传入一个interface类型的变量,可以表示任何值
*/
type Invoker interface {
	Call(interface{})
}

/*
这是一个结构体,其中它没有任何属性
*/
type Struct struct {
}

/*
定义这个结构体实现Invoker接口的函数
*/
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

/*
对于函数的声明而言,不能直接传递接口
而是要将函数定义为一个类型后,使用类型实现结构体
*/
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func main() {
	var invoker Invoker
	s := new(Struct)
	invoker = s
	invoker.Call([]string{"hi", "you"})

	/*
		定义了一个匿名函数,用于FuncCaller进行调用
	*/
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	invoker.Call("hello")

}
