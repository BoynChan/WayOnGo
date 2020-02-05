package main

import "fmt"

func main() {
	// 创建一个键为string类型,值为int类型的字典

	mapCreated := make(map[string]int)
	// 直接通过[KEY]=VALUE赋值
	mapCreated["Ten"] = 10
	mapCreated["One"] = 1

	fmt.Println(mapCreated["One"])
	fmt.Println(mapCreated["Ten"])
	//如果某个键不存在,会将其初始化为这个类型的初始值
	fmt.Println(mapCreated["Two"])

	for k, v := range mapCreated {
		fmt.Printf("%s %d\n", k, v)
	}

	//使用delete从map中删除一组键值对
	delete(mapCreated, "One")

	fmt.Println(mapCreated["One"])
}
