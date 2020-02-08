package main

import "fmt"

var eventByName = make(map[string][]func(interface{}))

type Actor string

// Actor的事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("Actor Event:", param)
}

// 全局事件处理函数
func GlobalEvent(param interface{}) {
	fmt.Println("Global Event:", param)
}

// 注册事件,需要提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	list := eventByName[name]
	list = append(list, callback)
	eventByName[name] = list
}

func CallEvent(name string, param interface{}) {
	funcList := eventByName[name]
	for _, v := range funcList {
		v(param)
	}
}

func main() {
	actor := new(Actor)
	//注册一个角色对应的事件和一个全局事件
	RegisterEvent("OnSkill", actor.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	//调用事件
	CallEvent("OnSkill", "hi")
}
