package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		time.Sleep(2 * time.Second)

		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			_, err := io.WriteString(conn, input.Text()+"\n")
			if err != nil {
				return // e.g., client disconnected
			}
			// time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Fatal(err)
		}
	}()
	for {

	}
}
