package runner

import (
	"fmt"
	"testing"
	"time"
)

/**
author:Boyn
date:2020/2/17
*/
var runner *Runner

func TestRunner(t *testing.T) {
	func1 := func(i int) {
		// 这会使runner超时
		time.Sleep(200 * time.Millisecond)
	}
	func2 := func(i int) {
		//do something
	}
	runner = New(100 * time.Millisecond)
	runner.Add(func1, func2)
	err := runner.Start()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}
