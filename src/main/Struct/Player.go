package main

import "fmt"

/*
定义一个player结构体,并对他进行初始化
*/

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

func main() {
	bob := new(Player)
	bob.HealthPoint = 15
	bob.Name = "Bob"
	bob.MagicPoint = 15
	fmt.Println(*bob)
}
