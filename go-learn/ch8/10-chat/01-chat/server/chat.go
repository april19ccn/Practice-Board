package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster() // 启动广播（广播纪元？hhh）
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // 启动协程，并发处理多个客户端
	}
}

// -------- broadcaster ---------------
type client chan<- string // an outgoing message channel  传出消息通道

var (
	entering = make(chan client) // 记录谁进入的通道，拿 消息通道 当key
	leaving  = make(chan client) // 记录谁退出的通道，拿 消息通道 当key
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients 客户端列表
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all 向所有人广播传入消息
			// clients' outgoing message channels. 客户的传出消息渠道。
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

func handleConn(conn net.Conn) {
	ch := make(chan string)   // outgoing client messages
	go clientWriter(conn, ch) // 通过消息通道将消息下发给客户端

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived" //为什么自身收不到这个消息，此时还没有加入到客户端列表
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
