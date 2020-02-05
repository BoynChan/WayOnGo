package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	//[1:2]这个就是切片的操作,注意切片是左闭右开的
	fmt.Println(a[1:2])

	var result [30]int
	for i := 0; i < 30; i++ {
		result[i] = i
	}
	//区间
	fmt.Println(result[10:15])
	//从20开始到结束
	fmt.Println(result[20:])
	//从开始到10
	fmt.Println(result[:10])

	//动态创建切片
	// 第一个参数是切片的元素类型,[]Type
	// 第二个参数size是为这个类型分配多少个元素
	// 第三个参数capacity是预分配空间,不影响size,但是可以降低多次分配空间造成的性能问题
	array1 := make([]int, 2)
	array2 := make([]int, 2, 10)

	fmt.Println(array1, array2)
	fmt.Println(len(array1), len(array2))

	//append用于为切片增加元素,如果空间不足,就会扩容
	array1 = append(array1, 1, 2, 3, 4)
	fmt.Println(array1)
	fmt.Println(len(array1))

	//在头部添加元素
	array1 = append([]int{5, 6}, array1...)
	fmt.Println(array1)
	//链式添加元素,以在第i个位置插入元素为例
	i := 2
	array1 = append(array1[:i], append([]int{10}, array1[i:]...)...)
	fmt.Print(array1)

	//对于切片的复制,如果目的切片比源切片要短,那么就会出现截断而不是扩容
	arrayCopy := make([]int, 5)
	copy(arrayCopy, array1)
	fmt.Println(arrayCopy)

}
