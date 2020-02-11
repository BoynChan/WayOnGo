package Counter

import (
	"testing"
)

func TestWordCounter_Write(t *testing.T) {
	w := new(WordCounter)
	write, _ := w.Write([]byte("a be cat dogs"))
	if write != 4 {
		t.Fail()
	}
}
