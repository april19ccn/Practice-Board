// 练习 8.2： 实现一个并发FTP服务器。
// 服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。

package main

import (
	"example/learn/ch8/02-example-clock/ex-2-ftp/utils"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 设置服务器端口
var port = flag.String("port", "8000", "port number")

type ftpServer struct {
	conn        net.Conn // 当前连接
	currentPath string   // 当前工作目录
	order       string   // 命令
	sync.Mutex           // 互斥锁
}

// 实现 io.Writer，用于从客户端读取命令
func (s *ftpServer) Write(p []byte) (n int, err error) {
	s.order = strings.ReplaceAll(string(p), "\n", "")
	s.order = strings.ReplaceAll(s.order, "\r", "")
	return len(p), nil
}

// 模仿控制台输出路径
func (s *ftpServer) Cmd() {
	str := time.Now().Format("2006-01-02 15:04:05 ") + "FTP SERVER: " + s.currentPath + " ❯ "
	if _, err := io.WriteString(s.conn, str); err != nil {
		log.Fatal(err)
	}
}

// 处理客户端发来的请求，通过io.Copy写入ftp的order属性
// 注意，这里的io.Copy是阻塞的，需要使用goroutines
func (s *ftpServer) Request() {
	defer s.conn.Close()

	if _, err := io.Copy(s, s.conn); err != nil {
		log.Fatal(err)
	}
}

// 响应客户端发来的请求，通过io.Copy写入ftp的order属性
// cd： 切换目录，通过 io.WriteString 返回错误信息
// ls： 列出目录内文件，通过 io.WriteString 返回目录内文件信息或错误
// get： 获取文件，通过 Get方法中 io.Copy 返回文件，io.Copy读取完文件就会结束，不会阻塞程序
func (s *ftpServer) Response() {
	defer s.conn.Close()

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
				if len(order) >= 2 {
					err = s.Get(order[1])
					if err != nil {
						response = err.Error() + "\n"
					}
				}
			case "send":
				// s.Send(order[1])
			case "close":
				s.Close()
			}

			if _, err := io.WriteString(s.conn, response); err != nil {
				log.Fatal(err)
			}
			s.Cmd()

			s.order = ""
			s.Unlock()
		}
	}

}

func (s *ftpServer) Close() {
	defer s.conn.Close()
	_, err := io.WriteString(s.conn, "close FTP SERVER\n")
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

func (s *ftpServer) Get(fileName string) error {
	if !filepath.IsAbs(fileName) {
		fileName = filepath.Join(s.currentPath, fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(s.conn, file)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *ftpServer) Send() {

}

// 启动服务器，并根据请求创建ftpServer（goroutine）
func Serve() {
	fmt.Println("FTP SERVER START: " + "localhost:" + *port)

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

		go CreateFTP(conn)
	}
}

// 创建ftpServer（goroutine）
// 每一个请求对应一个独立的ftpServer结构体和两个并发处理方法
func CreateFTP(conn net.Conn) {
	currentPath, _ := utils.GetWorkPath()

	ftp := &ftpServer{conn, currentPath, "", sync.Mutex{}}
	ftp.Cmd()

	// 由于8.2章还没有学到信道通信，所以两个协程是依赖ftpServer的order属性来进行通信
	// Response 在处理order时会上锁，防止Request修改，同时在处理期间本身也是要等服务器回应
	go ftp.Request()  // Request 会一直监听客户端的指令，必须单独运行一个协程
	go ftp.Response() // Response 会一直监听ftp的order属性，通过order属性的变化来响应客户端
}

func main() {
	flag.Parse()
	Serve()
	// ftp.Ls()
	// ftp.Cd("..")

	// err = ftp.Cd("./ch8/02-example-clock/ex-2-ftp/ftp.go")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// ftp.Ls()
}
