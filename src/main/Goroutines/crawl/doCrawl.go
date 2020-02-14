package crawl

import (
	"fmt"
	"strings"
)

/**
author:Boyn
date:2020/2/14
*/

//是否打印爬取细节
var detail = true

/*
使用爬虫的主函数
提供初始化根链接,防爬重复链接与按照BFS遍历链接的功能
*/
func DoCrawl2(rootUrl []string) {
	workList := make(chan []string)
	// n 指示workList中元素的个数
	n := 0
	n++
	go func() { workList <- rootUrl }()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				if detail {
					fmt.Printf("Seen:%s\n", link)
				}
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl2(link)
				}(link)
			}
		}
	}
}

/*
进行爬取的主函数,增加了深度的限制
*/
func DoCrawl3(rootUrl []string) {
	workList := make(chan Pages)
	// n 指示workList中元素的个数
	n := 0
	n++
	go func() { workList <- Pages{rootUrl, 1} }()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		pages := <-workList
		depth := pages.depth
		//如果url列表长度是0的话,说明发生错误或者爬到了对应的深度,于是就停止对这个链接进行爬取
		if len(pages.url) == 0 {
			continue
		}
		for _, link := range pages.url {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl3(link, depth)
				}(link)
			}
		}
	}
}

/*
判断是否为HTTP链接
*/
func isHttp(url string) bool {
	return strings.HasPrefix(url, "http")
}
