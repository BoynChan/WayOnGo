package Counter

import (
	"fmt"
	"testing"
)

func TestByteCounter_Write(t *testing.T) {
	var c ByteCounter
	write, _ := c.Write([]byte("hello"))
	if write != 5 {
		t.Error("长度不对")
	}
	c = 0
	var name = "Bob"
	fmt.Fprintf(&c, "hello, %s", name)
	if int(c) != 10 {
		t.Error("长度不对")
	}
}
