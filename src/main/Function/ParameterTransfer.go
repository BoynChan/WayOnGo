package main

import "fmt"

type Data struct {
	complex  []int      //切片类型
	instance InnerData  //结构体
	prt      *InnerData //结构体指针
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	//成员情况
	fmt.Printf("inFunc value : %+v\n", inFunc)
	//指针地址
	fmt.Printf("inFunc ptr : %p\n", &inFunc)
	return inFunc
}

/*
inFunc value : {complex:[1 2 3] instance:{a:5} prt:0xc00000a128}
inFunc ptr : 0xc00006a330
inFunc value : {complex:[1 2 3] instance:{a:5} prt:0xc00000a128}
inFunc ptr : 0xc00006a3c0
inFunc value : {complex:[1 2 3] instance:{a:5} prt:0xc00000a128}
inFunc ptr : 0xc00006a390

可以看到所有Data的指针地址都发生了变化,说明我们直接传递Data是进行复制传递,同时返回结果也是复制的
只有指针,切片和map在传递过程中才会进行地址的传递
*/
func main() {
	inFunc := Data{
		complex:  []int{1, 2, 3},
		instance: InnerData{a: 5},
		prt:      &InnerData{a: 1},
	}
	//成员情况
	fmt.Printf("inFunc value : %+v\n", inFunc)
	//指针地址
	fmt.Printf("inFunc ptr : %p\n", &inFunc)
	out := passByValue(inFunc)

	//成员情况
	fmt.Printf("inFunc value : %+v\n", out)
	//指针地址
	fmt.Printf("inFunc ptr : %p\n", &out)

}
