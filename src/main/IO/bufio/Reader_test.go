package bufio

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

// Author:Boyn
// Date:2020/2/19
var header = `GET http://download.microtool.de:80/somedata.exe
		Host: download.microtool.de
		Accept:*/*
		Pragma: no-cache
		Cache-Control: no-cache
		Referer: http://download.microtool.de/
		User-Agent:Mozilla/4.04[en](Win95;I;Nav)
		Range:bytes=554554-`

func TestReadSlice(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("https://studygolang.com.\nI am boyn"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("line:%s\n", line)
	line, _ = reader.ReadSlice('\n')
	fmt.Printf("line:%s\n", line)
}

func TestReadBytes(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("https://studygolang.com\nI am boyn\n"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("line:%s\n", line)
	line, _ = reader.ReadBytes('\n')
	fmt.Printf("line:%s\n", line)
}

func TestReadString(t *testing.T) {

	reader := bufio.NewReader(strings.NewReader(header))
	bytes, err := reader.ReadString('\n')
	if err != nil {
		//fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

// 使用Scanner进行行数统计
func TestScannerCountLines(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(header))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}
