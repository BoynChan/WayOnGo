package io

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// Author:Boyn
// Date:2020/2/19

func TestReadFromStdIn(t *testing.T) {
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func TestReadAtFromStr(t *testing.T) {
	data, n, err := ReadAt(strings.NewReader("Hi I'm Boyn"), 11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s %d\n", data, n)
}

func TestReadFromStr(t *testing.T) {
	data, err := ReadFrom(strings.NewReader("Hi I'm Boyn"), 11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func TestReadFromFunction(t *testing.T) {
	open, err := os.Open("F:\\Code\\Go\\LearningGo\\.gitignore")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, _, err := ReadFromIO(open)
	fmt.Println(string(content))
}
