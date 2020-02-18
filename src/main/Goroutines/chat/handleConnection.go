package chat

import (
	"bufio"
	"fmt"
	"net"
)

/**
author:Boyn
date:2020/2/15
*/

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	name, err := Login(ch, conn)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String() + " login fail")
		conn.Close()
	}
	messages <- name + " has arrived"
	cli := client{ip: who, name: name, channel: ch, conn: conn}
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- name + ":" + input.Text()
	}
	leaving <- cli
	messages <- name + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintln(conn, msg)
	}
}

func Login(ch chan string, conn net.Conn) (string, error) {
	ch <- "Please enter your name:\n"
	input := bufio.NewScanner(conn)
	input.Scan()
	name := input.Text()
	ch <- "Welcome!" + name
	return name, nil
}
