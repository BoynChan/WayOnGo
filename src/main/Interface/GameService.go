package main

import "fmt"

type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

func (g *Logger) Log(s string) {
	fmt.Println("[Log]:", s)
}

//对于GameService来说,如果他需要实现Service接口
//可以实现两个方法,但是也可以选择像下面这样,只是实现一个方法,另外方法的实现靠与别的struct组合完成
//这种方法与面向对象十分不一样,他不是以继承的方式来实现类复用,而是在很多的时候使用了组合
type GameService struct {
	Logger
}

func (g *GameService) Start() {
	fmt.Println("Game start")
}

func main() {

}
