package io

import (
	"bytes"
	"fmt"
	"io"
)

// Author:Boyn
// Date:2020/2/19

// 将io.Reader和int作为参数,表示从Reader中读取num个字符
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func ReadAt(reader io.ReaderAt, num int) ([]byte, int, error) {
	p := make([]byte, num)
	n, err := reader.ReadAt(p, 0)
	if err != nil {
		return p, 0, err
	}
	return p, n, nil
}

func ReadFromIO(reader io.Reader) ([]byte, int, error) {
	var buf bytes.Buffer
	// 使用Buffer中的ReadFrom函数来读取Reader中的内容
	n, err := buf.ReadFrom(reader)
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes(), int(n), err
}
