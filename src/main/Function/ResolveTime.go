package main

import "fmt"

const (
	SecondPerMinute = 60
	SecondPerHour   = SecondPerMinute * 60
	SecondPerDay    = SecondPerHour * 24
)

/*
将秒转换为具体的时间
*/
func main() {
	fmt.Println(resolveTime(3600 * 24))
}

func resolveTime(second int) (minutes, hours, days int) {
	return second / SecondPerMinute, second / SecondPerHour, second / SecondPerDay
}
