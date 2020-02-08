package main

import "fmt"

func main() {
	str := "hello world"

	//闭包特性,将引用传递了进去,并对引用的对象进行了修改
	// hello world add:0xc0000561e0
	// hello add:0xc0000561e0
	// 从中可以看到,内存地址并没有改变,因为闭包特性,不能修改传入的地址
	fc := func() {
		str = "hello"
	}
	fmt.Printf("%s add:%p\n", str, &str)
	fc()
	fmt.Printf("%s add:%p\n", str, &str)

	player := genPlayer("Bob")
	fmt.Println(player())
}

// 通过闭包,我们可以创建一个类似于工厂方法的实体
// 在python中,可以被称为生成器
func genPlayer(name string) func() (string, int) {
	//设置血量为50
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}
