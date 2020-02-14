package crawl

import (
	"fmt"
	"main/Net"
)

/**
并发爬虫
author:Boyn
date:2020/2/14
*/

/*
这个被运行的并行函数没有对并发量做任何设置
运行起来会使得go进程超出系统的资源限制
所以我们要对其进行一定的限制
*/
func crawl1(url string) []string {
	if !isHttp(url) {
		return make([]string, 0)
	}
	if detail {
		fmt.Printf("Crawl:%s\n", url)
	}
	links, err := Net.FindLinks(url)
	if err != nil {
		return make([]string, 0)
	}
	return links
}
