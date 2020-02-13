package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

/**
对于比较复杂的数据结构,我们可以使用Interface接口的实现
来让其能够进行排序
author:Boyn
date:2020/2/13
*/

/*
表示唱片的结构体
*/
type Track struct {
	Title  string        //唱片名称
	Artist string        //唱片作者
	Album  string        //所属专辑
	Year   int           //发行年份
	Length time.Duration //唱片时长
}

/*
以作者名字进行排序
*/
type byArtist []*Track

func (b byArtist) Len() int {
	return len(b)
}

func (b byArtist) Less(i, j int) bool {
	return b[i].Artist < b[j].Artist
}

func (b byArtist) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

/*
以年份进行排序
*/
type byYear []*Track

func (b byYear) Len() int {
	return len(b)
}

func (b byYear) Less(i, j int) bool {
	return b[i].Year < b[j].Year
}

func (b byYear) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func PrintTracks(track []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range track {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

/**
以作者名字作为键进行排序
*/
func SortByArtist(track []*Track) {
	sort.Sort(byArtist(track))
}

/*
以年份作为键进行排序
*/
func SortByYear(track []*Track) {
	sort.Sort(byYear(track))
}

/*
用于将文本形式的时间转换为time.Duration形式
*/
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
