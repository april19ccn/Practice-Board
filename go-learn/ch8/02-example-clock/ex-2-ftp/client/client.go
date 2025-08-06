package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// 发送指令给 ftp 服务器
func SendCommand(conn net.Conn) {
	time.Sleep(2 * time.Second)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		_, err := io.WriteString(conn, input.Text()+"\n")
		if err != nil {
			return
		}
	}
}

// 展示 ftp 服务器返回的结果
func ShowResult(conn net.Conn) {
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go SendCommand(conn)

	go ShowResult(conn)

	for {

	}
}
