package main

import (
	"flag"
	"fmt"
)

//命令行参数为mode,默认值为空字符串,mode的类型是*string
var mode = flag.String("mode", "", "process mode")

func main() {
	// go run .\src\main\CommandLineInput.go --mode=Hi
	flag.Parse()
	fmt.Println(*mode)
}
