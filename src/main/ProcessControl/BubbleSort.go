package main

import "fmt"

func main() {
	array := []int{5, 2, 7, 43, 9, 2, 76}
	sort(array)
	fmt.Println(array)
}
func sort(array []int) {
	var isSwap bool
	for i := 0; i < len(array); i++ {
		isSwap = false
		for j := i; j < len(array); j++ {
			if array[j] < array[i] {
				isSwap = true
				array[j], array[i] = array[i], array[j]
			}
		}
		if !isSwap {
			break
		}
	}
}
