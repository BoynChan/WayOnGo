package Reader

import (
	"fmt"
	"testing"
)

func TestStringReader(t *testing.T) {
	r := NewStrReader("Hi!I am boyn")
	b := make([]byte, 6)
	read, _ := r.Read(b)
	fmt.Println(string(b))
	if read != 6 {
		t.Fail()
	}
	read, _ = r.Read(b)
	fmt.Println(string(b))
	if read != 6 {
		t.Fail()
	}
}
