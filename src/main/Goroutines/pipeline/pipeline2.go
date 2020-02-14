package pipeline

import (
	"fmt"
)

/**
第二个流水线使用到了通道的可迭代性,在通道被关闭的时候及时关闭
author:Boyn
date:2020/2/14
*/

func pipeline2() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		//只有10个数字
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for num := range naturals {
			squares <- num * num
		}
		close(squares)
	}()

	for {
		for num := range squares {
			fmt.Println(num)
		}
		break
	}
}
