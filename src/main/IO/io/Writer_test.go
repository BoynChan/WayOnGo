package io

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// Author:Boyn
// Date:2020/2/19

func TestWriteAt(t *testing.T) {
	file, err := os.Create("F:\\Code\\Go\\LearningGo\\TestWriteAtFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 中文一个字占据3个字节,先写入我是谁? 然后在偏移量为6的位置写入Boyn, 那么结果就是Boyn
	_, _ = WriteAt(file, 0, []byte("我是谁?"))
	_, _ = WriteAt(file, 6, []byte("Boyn"))
}

func TestWriteTo(t *testing.T) {
	reader := bytes.NewReader([]byte("Hi I'm Boyn\n"))
	WriteTo(reader, os.Stdout)
}
