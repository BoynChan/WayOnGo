package bytes

import (
	"bytes"
	"fmt"
	"testing"
)

// bytes下的包大部分操作与strings相同,所以仅演示部分
// Author:Boyn
// Date:2020/2/20

func TestBuffer(t *testing.T) {
	// Buffer封装了[]byte并提供了很多方便的操作
	var buf bytes.Buffer
	buf.Write([]byte{'a', 'b', 'c'})
	fmt.Println(buf.Len())
	fmt.Println(buf.String())
}

func TestContains(t *testing.T) {
	var buf bytes.Buffer
	buf.Write([]byte{'a', 'b', 'c'})
	var buf2 bytes.Buffer
	buf2.Write([]byte{'a'})
	fmt.Println(bytes.Contains(buf.Bytes(), buf2.Bytes()))
}

func TestCount(t *testing.T) {
	var buf bytes.Buffer
	buf.Write([]byte{'a', 'b', 'c', 'a'})
	var buf2 bytes.Buffer
	buf2.Write([]byte{'a'})
	fmt.Println(bytes.Count(buf.Bytes(), buf2.Bytes()))
}
