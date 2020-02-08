package main

import "fmt"

type Walkable struct{}

func (w *Walkable) Walk() {
	fmt.Println("Walk")
}

type Flyable struct {
}

func (f *Flyable) Fly() {
	fmt.Println("Fly")
}

//通过内嵌的方式模拟继承
//但是这并不是完全的继承,因为没有完全实现面向对象中的多态
type Human struct {
	Walkable
}

type Bird struct {
	Walkable
	Flyable
}

func main() {
	b := new(Bird)
	b.Fly()
	b.Walk()

	h := new(Human)
	h.Walk()
}
