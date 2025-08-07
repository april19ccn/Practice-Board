package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type ftpClient struct {
	conn        net.Conn // 当前连接
	defaultPath string   // 默认下载地址
	order       string   // 当前命令
}

func (ftp *ftpClient) Write(p []byte) (n int, err error) {
	order := strings.Split(ftp.order, " ")

	fmt.Println(order)

	if len(order) >= 2 && order[0] == "get" && !strings.Contains(string(p), "FTP SERVER") {
		HandleGet(p, order)
	} else {
		fmt.Print(string(p))
	}
	return len(p), nil
}

// 发送指令给 ftp 服务器
func (ftp *ftpClient) SendCommand() {
	time.Sleep(2 * time.Second)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		_, err := io.WriteString(ftp.conn, input.Text()+"\n")
		ftp.order = input.Text()
		if err != nil {
			return
		}
	}
}

// 处理 ftp 服务器返回的结果
func (ftp *ftpClient) HandleResult() {
	if _, err := io.Copy(ftp, ftp.conn); err != nil {
		log.Fatal(err)
	}
}

// 处理 get 命令
func HandleGet(t []byte, target []string) {
	// 如果文件不存在：创建新文件
	// 如果文件已存在：清空文件内容（截断为0字节）
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, bytes.NewReader(t))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ftp := &ftpClient{conn, "", ""}

	go ftp.SendCommand()

	go ftp.HandleResult()

	for {

	}
}
