package main

import (
	"fmt"
	"log"
)

//Author: Boyn
//Date: 2020/4/2

func main() {
	client, err := DialHelloService("tcp", "localhost:12346")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
