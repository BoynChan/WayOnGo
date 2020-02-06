package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
聊天机器人
*/
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Input your name:")
	input, err := inputReader.ReadString('\n')
	exitByErrorIfExist(err)
	name := input[:len(input)-1]
	fmt.Printf("Hi %s! Talk with me\n", name)

	for {
		fmt.Printf("Speak:")
		input, err := inputReader.ReadString('\n')
		exitByErrorIfExist(err)
		input = input[:len(input)-1]
		switch input {
		case "":
			continue
		case "bye", "goodbye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("You just say ", input)

		}
	}
}

func exitByErrorIfExist(err error) {
	if err != nil {
		fmt.Println("错误发生:", err)
		os.Exit(1)
	}
}
