package main

import (
	"fmt"
	"strings"
)

type StringHandlerChain []func(string) string

/*
链式编程
*/
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}
func main() {
	list := []string{
		"go language",
		"go player",
		"go printer",
		"go runtime",
	}
	chain := StringHandlerChain{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	StringProcess(list, chain)
	fmt.Println(list)
}

func StringProcess(list []string, chain StringHandlerChain) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}
