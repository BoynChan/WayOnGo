package main

import "testing"

/**
author:Boyn
date:2020/2/14
*/

func TestMultiListen(t *testing.T) {
	//这个测试启动了4个服务器内容,分别指代重庆时间,东京时间,堪培拉时间,伦敦时间
	go Listen(8990, "Asia/Chongqing")
	go Listen(8991, "Asia/Tokyo")
	go Listen(8992, "Australia/Canberra")
	go Listen(8993, "Europe/London")
	select {}
}
