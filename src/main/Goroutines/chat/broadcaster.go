package chat

/**
author:Boyn
date:2020/2/15
*/

type client chan<- string //定义一个输出channel,表示向客户端中输出消息

var (
	entering = make(chan client) // 监控客户端进入的消息
	leaving  = make(chan client) // 监控客户端离开的消息
	messages = make(chan string) // 掌握所有客户端发出的消息
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
