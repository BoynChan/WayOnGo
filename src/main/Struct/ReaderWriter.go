package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	ReadTest()
	ReadByteTest()
	ReadBytesTest()
	ReadLineTest()
}

func ReadTest() {
	data := []byte("学习Go语言")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
}

func ReadByteTest() {
	data := []byte("Go语言")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	b, err := r.ReadByte()
	fmt.Println(string(b), err)
}

func ReadBytesTest() {
	//ReadBytes会一直读取数据直到遇到一个分隔符为止
	delim := byte(',')
	data := []byte("Learn Go Language, and do it")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	readBytes, e := r.ReadBytes(delim)
	fmt.Println(string(readBytes), e)
}

func ReadLineTest() {
	data := []byte("Go is beautiful.\r\n I Like it")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)
}
