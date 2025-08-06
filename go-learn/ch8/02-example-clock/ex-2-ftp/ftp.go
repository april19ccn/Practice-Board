// 练习 8.2： 实现一个并发FTP服务器。
// 服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。

package main

import (
	"fmt"
	"log"
	"os"
)

type ftpServer struct {
	currentPath string
}

// 启动服务
func (s *ftpServer) Serve() {

}

func (s *ftpServer) Close() {

}

func (s *ftpServer) Cd(path string) error {

	return nil
}

func (s *ftpServer) Ls() []os.DirEntry {
	files, err := os.ReadDir(s.currentPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("文件夹: %s\n", file.Name())
		} else {
			fmt.Printf("文件: %s\n", file.Name())
		}
	}
	return files
}

func (s *ftpServer) Get() {

}

func (s *ftpServer) Send() {

}

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ftp := &ftpServer{currentPath}
	// ftp.Serve()
	ftp.Ls()

}
