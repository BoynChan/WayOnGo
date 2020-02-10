package Net

import (
	"fmt"
	"testing"
)

func TestFindLinks1(t *testing.T) {
	strings, _ := FindLinks("http://gopl.io")
	for _, url := range strings {
		fmt.Println(url)
	}
}

func TestOutline(t *testing.T) {
	node, _ := getPageRootNode("http://gopl.io")
	Outline(nil, node)
}

func TestCountWordsAndImages(t *testing.T) {
	words, images, _ := CountWordsAndImages("http://gopl.io")
	fmt.Println(words, images)
}
