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

	var reply String
	err = client.Hello(String{Value:"Hello"}, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.Value)
}
