// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "连接错误: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close()

// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintf(os.Stderr, "读取错误: %v\n", err)
// 	}
// }

// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
