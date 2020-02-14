package crawl

import (
	"fmt"
	"main/Net"
)

/**
author:Boyn
date:2020/2/14
*/

func crawl3(url string, depth int) Pages {
	//如果深度大于3,那么直接返回空列表
	if depth > 3 {
		return Pages{url: make([]string, 0), depth: depth + 1}
	}
	if detail {
		fmt.Printf("Depth:%d Crawl:%s\n", depth, url)
	}
	//向tokens中放入一个元素,提供占位的功能
	tokens <- struct{}{}
	//函数结束时将tokens中的元素出队
	defer func() { <-tokens }()
	links, err := Net.FindLinks(url)
	//在爬取结束后,先对非HTTP链接进行过滤再返回
	links = filterHttp(links)
	if err != nil {
		return Pages{url: make([]string, 0), depth: depth + 1}
	}
	return Pages{url: links, depth: depth + 1}
}

//过滤一个列表,筛选出http的链接
func filterHttp(url []string) []string {
	supplierChannel := make(chan string)
	filterChannel := make(chan string)
	go func() {
		for _, v := range url {
			supplierChannel <- v
		}
		close(supplierChannel)
	}()
	go func() {
		for url := range supplierChannel {
			if isHttp(url) {
				filterChannel <- url
			}
		}
		close(filterChannel)
	}()
	result := make([]string, 0)
	for url := range filterChannel {
		result = append(result, url)
	}
	return result
}
