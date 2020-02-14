package pipeline

import "fmt"

/**
将channel进行拆分,分出3个函数分别进行流水线操作
author:Boyn
date:2020/2/14
*/

func pipeline3() {
	counterChannel := make(chan int)
	squarerChannel := make(chan int)
	go counter(counterChannel)
	go squarer(squarerChannel, counterChannel)
	printer(squarerChannel)
}

/*
计数器,通道只会用于输出
*/
func counter(out chan int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

/*
平方器,有一个用于输入的管道和输出的管道
*/
func squarer(out, in chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

/*
打印器,只有一个用于输入的管道
*/
func printer(in chan int) {
	for num := range in {
		fmt.Println(num)
	}
}
