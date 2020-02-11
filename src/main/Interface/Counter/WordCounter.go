package Counter

import (
	"bufio"
	"strings"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	split := strings.Split(string(p), " ")
	bufio.ScanWords()
	*w += WordCounter(len(split))
	return len(split), nil
}
