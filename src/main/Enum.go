package main

import "fmt"

//以iota作为枚举值
//并且可以对iota做数值计算,使得他能够作为一个数值的生成器
type Weapon int

const (
	Arrow Weapon = iota * 2 //plus(iota) 不能使用函数对其调用,因为常量值要在编译期间确定
	Shuriken
	SniperRifle
	Rifle
	Blower
)

func plus(a int) int {
	return a + 1
}

//使用String()函数作为类似对象的toString函数将枚举类型转换为字符串
func (w Weapon) String() string {
	switch w {
	case Arrow:
		return "Arrow"
	case Blower:
		return "Blower"
	case Rifle:
		return "Rifle"
	}
	return ""
}

func main() {
	var weapon Weapon = Blower
	fmt.Printf("%s %d", weapon, weapon)
}
