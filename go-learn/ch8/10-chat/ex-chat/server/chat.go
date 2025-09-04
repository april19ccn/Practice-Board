package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
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

type clientInfo struct {
	name string
	ch   client
}

var (
	entering = make(chan clientInfo) // 记录谁进入的通道，拿 消息通道 当key
	leaving  = make(chan clientInfo) // 记录谁退出的通道，拿 消息通道 当key
	messages = make(chan string, 5)  // all incoming client messages
)

func _getClients(current string, clients map[string]client) string {
	result := "\n ---------- 🪐 Current Clients ------------ \n"

	for k := range clients {
		if k == current {
			result += k + " ⭐ " + "\n"
			continue
		}
		result += k + "\n"
	}

	result += "----------------\n"
	return result
}

func broadcaster() {
	var clients = make(map[string]client) // all connected clients 客户端列表
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all 向所有人广播传入消息
			// clients' outgoing message channels. 客户的传出消息渠道。
			for _, ch := range clients {
				ch <- msg
			}

		case enter := <-entering:
			clients[enter.name] = enter.ch
			enter.ch <- _getClients(enter.name, clients)

		case leave := <-leaving:
			if ch, exists := clients[leave.name]; exists {
				delete(clients, leave.name)
				close(ch)
			}
		}
	}
}

const AUTO_CLOSE_TIME = 1000

func handleConn(conn net.Conn) {
	ch := make(chan string)   // outgoing client messages
	go clientWriter(conn, ch) // 通过消息通道将消息下发给客户端

	input := bufio.NewScanner(conn)

	// 设置name
	who := conn.RemoteAddr().String()
	ch <- "What is your name?"
	for input.Scan() {
		who = input.Text()
		break
	}

	messages <- who + " has arrived" //为什么自身收不到这个消息，此时还没有加入到客户端列表
	entering <- clientInfo{who, ch}

	var once sync.Once
	closeClient := func() {
		leaving <- clientInfo{who, ch}
		messages <- who + " has left"
		conn.Close()
	}

	timer := time.NewTimer(AUTO_CLOSE_TIME * time.Second)
	go func() {
		<-timer.C
		once.Do(closeClient)
	}()

	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(AUTO_CLOSE_TIME * time.Second)
	}

	once.Do(closeClient)
	timer.Stop()
	fmt.Println(who + " go handleConn 结束")
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
