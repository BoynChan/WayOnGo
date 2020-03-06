package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

// Author:Boyn
// Date:2020/3/6

func main() {
	test := &Student{
		Name:   "Bob",
		Male:   true,
		Scores: []int32{98, 95, 89},
	}
	//fmt.Println(test.String())
	data, _ := proto.Marshal(test)
	newTest := &Student{}
	_ = proto.UnmarshalMerge(data, newTest)
	fmt.Println(newTest.Scores)
}
