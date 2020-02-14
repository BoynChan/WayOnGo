package crawl

import (
	"fmt"
	"main/Net"
)

/**
author:Boyn
date:2020/2/14
*/

//tokens可以理解为一个信号量,只允许最多20个并发的网络请求
var tokens = make(chan struct{}, concurrentRequests)

const concurrentRequests int = 20

func crawl2(url string) []string {
	if !isHttp(url) {
		return make([]string, 0)
	}
	if detail {
		fmt.Printf("Crawl:%s\n", url)
	}
	//向tokens中放入一个元素,提供占位的功能
	tokens <- struct{}{}
	//函数结束时将tokens中的元素出队
	defer func() { <-tokens }()
	links, err := Net.FindLinks(url)
	if err != nil {
		return make([]string, 0)
	}
	return links
}
