package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 6, 7, 8, 10}
	fmt.Println(BinarySearch(array, 10))
}

/*
接收一个array数组与目标值,返回目标值在数组中的索引
*/
func BinarySearch(array []int, target int) int {
	return binarySearch(array, 0, len(array), target)
}

func binarySearch(array []int, l, h, target int) int {
	mid := -1
	for l <= h {
		mid = l + (h-l)/2
		if array[mid] == target {
			break
		} else if array[mid] < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	if array[mid] == target {
		return mid
	} else {
		return -1
	}
}
