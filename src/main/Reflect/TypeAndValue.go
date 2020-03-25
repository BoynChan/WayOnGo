package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// Author:Boyn
// Date:2020/3/23

type User struct {
}

func main() {
	//TypeOf接收任意interface{}类型,并以Type形式返回其动态类型
	t := reflect.TypeOf(1)
	fmt.Println(t)
	t = reflect.TypeOf(User{})
	fmt.Println(t)
	t = reflect.TypeOf(&User{})
	fmt.Println(t)

	//TypeOf返回的是动态类型,所以如果传入的是接口,会返回接口对应的实现结构体
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	// 也可以直接用%T参数获取类型
	fmt.Printf("%T\n", w)

	//ValueOf可以返回某变量对应的值,下面两种写法相同
	// 并且在转换到Value类型后,可以用Kind来获取其属于的动态类型
	fmt.Println(reflect.ValueOf(1))
	fmt.Printf("%v\n", 1)

}
