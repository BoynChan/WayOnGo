package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Println(time.Since(start))
}
