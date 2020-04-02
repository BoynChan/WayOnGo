package model

import "fmt"

//Author: Boyn
//Date: 2020/3/29

type Poster interface {
	postArticle(p Post) error
}

type HexoPoster struct {
	url string
}

func (h *HexoPoster) postArticle(p Post) error {
	fmt.Println("Posted Article to Hexo on " + h.url)
	return nil
}

type HugoPoster struct {
	url string
}

func (h *HugoPoster) postArticle(p Post) error {
	fmt.Println("Posted Article to Hugo on " + h.url)
	return nil
}

type Post struct {
	title   string
	content string
	User
	Poster
}

func (p *Post) post() {
	_ = p.Poster.postArticle(*p)
}

func (p *Post) details() {
	fmt.Printf("Title: %s\nContent: %s\n User: %s\n", p.title, p.content, p.Username)
}
