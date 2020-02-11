package Reader

import (
	"fmt"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	reader := strings.NewReader("Hi!I am boyn")
	lR := NewLimitReader(reader, 10)
	b := make([]byte, 6)
	read, _ := lR.Read(b)
	fmt.Println(read)
	if read != 6 {
		t.Fail()
	}
	read, _ = lR.Read(b)
	fmt.Println(read)
	if read != 4 {
		t.Fail()
	}
}
