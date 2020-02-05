package main

import "fmt"

const name = "博言"

//使用iota来声明一系列递增的常量
type WeekDay int

const (
	Sunday WeekDay = iota
	MonDay
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	const pi = 3.1415926
	fmt.Println(pi)
	fmt.Println(name)

	fmt.Println(Sunday)
	fmt.Println(Friday)

}
