package echo1

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

/**
一个简单的echo服务器
author:Boyn
date:2020/2/14
*/

func Listen() {
	listener, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text())
	}
}

func echo(c net.Conn, shout string) {
	fmt.Fprintf(c, fmt.Sprintf("%s\n", strings.ToUpper(shout)))
}
