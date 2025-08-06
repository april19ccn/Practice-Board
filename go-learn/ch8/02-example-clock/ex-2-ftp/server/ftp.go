// 练习 8.2： 实现一个并发FTP服务器。
// 服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type ftpServer struct {
	currentPath string
	order       string
	sync.Mutex
}

func (s *ftpServer) Write(p []byte) (n int, err error) {
	s.order = strings.ReplaceAll(string(p), "\n", "")
	s.order = strings.ReplaceAll(s.order, "\r", "")
	return len(p), nil
}

// 启动服务
var port = flag.String("port", "8000", "port number")

func (s *ftpServer) Cmd(conn net.Conn) {
	str := time.Now().Format("2006-01-02 15:04:05 ") + "FTP SERVER: " + s.currentPath + " ❯ "
	if _, err := io.WriteString(conn, str); err != nil {
		log.Fatal(err)
	}
}

func (s *ftpServer) Serve() {
	flag.Parse()

	fmt.Println("FTP SERVER START: " + "localhost:" + *port)
	fmt.Println("FTP SERVER: ", s.currentPath)
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		s.Cmd(conn)

		go s.Request(conn)
		go s.Response(conn)
	}
}

func (s *ftpServer) Request(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(s, conn); err != nil {
		log.Fatal(err)
	}
}

func (s *ftpServer) Response(conn net.Conn) {
	defer conn.Close()

	for {
		if s.order != "" {
			s.Lock()
			order := strings.Split(s.order, " ")

			response := ""
			err := error(nil)

			switch order[0] {
			case "cd":
				if len(order) >= 2 {
					_, err = s.Cd(order[1])
					if err != nil {
						response = err.Error() + "\n"
					}
				}
			case "ls":
				response, err = s.Ls()
				if err != nil {
					response = err.Error()
				}
			case "get":
				// s.Get(order[1])
			case "send":
				// s.Send(order[1])
			case "close":
				s.Close(conn)
			}

			if _, err := io.WriteString(conn, response); err != nil {
				log.Fatal(err)
			}
			s.Cmd(conn)

			s.order = ""
			s.Unlock()
		}
	}

}

func (s *ftpServer) Close(conn net.Conn) {
	defer conn.Close()
	_, err := io.WriteString(conn, "close FTP SERVER\n")
	if err != nil {
		log.Fatal(err)
	}
}

// test 的时候 在测试文件是否能检测出这不是目录，可以先创建临时文件，然后删掉
func (s *ftpServer) Cd(path string) (string, error) {
	if !filepath.IsAbs(path) {
		path = filepath.Join(s.currentPath, path)
	}

	fmt.Println(path)

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("directory does not exist")
		}
		return "", err
	}
	// fmt.Println(info) // fmt.Println(info) 打印的是 os.FileInfo 的具体实现结构（Windows 上是 *os.fileStat），包含文件/目录的完整元数据。
	if !info.IsDir() {
		return "", fmt.Errorf("not a directory")
	}

	s.currentPath = path
	// fmt.Println("FTP SERVER: ", s.currentPath)
	return "FTP SERVER: " + s.currentPath + "\n", nil
}

func (s *ftpServer) Ls() (string, error) {
	files, err := os.ReadDir(s.currentPath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// lsResult := fmt.Sprint("FTP SERVER: " + s.currentPath + "\n")
	lsResult := fmt.Sprintf("%-5s %-25s %-15s %s\n", "💎", "ModTime", "Size", "Name")
	lsResult += fmt.Sprintf("%-5s %-25s %-15s %s\n", "------", "-------------------------", "---------------", "------")

	for _, file := range files {
		info, _ := file.Info()
		if file.IsDir() {
			lsResult += fmt.Sprintf("%-5s %-25s %-15s %s\n", "📁", info.ModTime().Format("2006-01-02 15:04:05"), "", file.Name())
		} else {
			lsResult += fmt.Sprintf("%-5s %-25s %-15d %s\n", "📄", info.ModTime().Format("2006-01-02 15:04:05"), info.Size(), file.Name())
		}
	}

	fmt.Println(lsResult)

	return lsResult, nil
}

func (s *ftpServer) Get() {

}

func (s *ftpServer) Send() {

}

func main() {
	// 获取当前服务器工作目录
	// currentPath, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 获取home目录
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	ftp := &ftpServer{u.HomeDir, "", sync.Mutex{}}
	ftp.Serve()
	// ftp.Ls()
	// ftp.Cd("..")

	// err = ftp.Cd("./ch8/02-example-clock/ex-2-ftp/ftp.go")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// ftp.Ls()
}
