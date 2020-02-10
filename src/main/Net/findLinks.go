package Net

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func FindLinks(url string) ([]string, error) {
	node, err := getPageRootNode(url)
	if err != nil {
		return nil, err
	}
	//递归遍历文本,将遇到的结点输出到link(一个slice)中
	return visit(nil, node), err

}

func getPageRootNode(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}

	//解析文本
	node, e := html.Parse(resp.Body)
	if e != nil {
		return nil, fmt.Errorf("parsing %s as HTML : %v", url, err)
	}
	return node, nil
}

/*
遍历所有结点,并打印出结点的层次结构
*/
func Outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Outline(stack, c)
	}
}

/*
遍历所有的结点,并将结点为超链接(a)的放入列表中
*/
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/*
遍历所有结点,并返回其中文本的词组数量和图片数量
*/
func CountWordsAndImages(url string) (words, images int, err error) {
	node, err := getPageRootNode(url)
	if err != nil {
		return
	}
	words, images = countWordsAndImages(node)
	return
}

/*
在每个结点中,如果是img结点,那就计数
如果是文本类的结点,就将其文本进行分词,计数
递归,相加
*/
func countWordsAndImages(n *html.Node) (words, images int) {
	words, images = 0, 0
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wordRecursive, imageRecursive := countWordsAndImages(c)
		images += imageRecursive
		words += wordRecursive
	}
	return
}

/*
接收一个HTML结点数以及任意数量的标签名,返回跟这些便签名匹配的所有元素
*/
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var result []*html.Node
	if doc.Type == html.ElementNode && contains(doc.Data, name) {
		result = append(result, doc)
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, ElementsByTagName(c, name...)...)
	}
	return result
}

func contains(name string, list []string) bool {
	for _, v := range list {
		if name == v {
			return true
		}
	}
	return false
}
