// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var isAccept = flag.Bool("a", false, "是否接收服务器数据")

// !+
func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		if *isAccept {
			io.Copy(os.Stdout, conn) // 客户端收数据
			log.Fatal("done")
		} else {
			reader := bufio.NewReader(conn)
			for i := 0; i < 5; i++ {
				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				fmt.Print(line)
			}

			// 模拟阻塞
			select {}

			// conn.(*net.TCPConn).CloseRead() // 服务器的 bufio.NewScanner(conn) 会认为是异常连接而结束读取
		}
		done <- struct{}{} // signal the main goroutine
	}()

	mustCopy(conn, os.Stdin) // 客户端写数据

	conn.Close()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
