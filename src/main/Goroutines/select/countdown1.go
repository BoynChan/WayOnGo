package _select

import (
	"fmt"
	"time"
)

/**
进行火箭发射的倒计时
author:Boyn
date:2020/2/15
*/

func countdown1() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second) //Tick函数返回一个阻塞channel,周期性地发送一个时间
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Fire.")
}
