package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

type city struct {
	name string
	url  string
	data string
}

func (c *city) Write(p []byte) (n int, err error) {
	// 去除所有换行符和回车符
	c.data = strings.ReplaceAll(string(p), "\n", "")
	c.data = strings.ReplaceAll(c.data, "\r", "")
	return len(p), nil
}

// Shanghai=localhost:8010
// Beijing=localhost:8020
// Xinjiang=localhost:8030
func main() {
	cityData := make([]city, 0, 8)

	for _, data := range os.Args[1:] {
		fmt.Println("访问：" + data)
		t := strings.Split(data, "=")
		cityData = append(cityData, city{name: t[0], url: t[1]})
		go getTime(&cityData[len(cityData)-1])
	}

	for {
		printTime(cityData)
	}
}

// 定义清屏函数（适配 PowerShell）
func clearScreen() {
	// 调用 PowerShell 的清屏命令
	cmd := exec.Command("powershell", "-Command", "Clear-Host")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printTime(citys []city) {
	// 使用 ANSI 转义序列操作终端的指令
	// \033[H：将光标移动到终端的左上角（"Home" 位置）
	// \033[2J：清除终端屏幕上的所有内容
	// powershell 可能不支持 ANSI 转义序列操作终端的指令
	// fmt.Print("\033[H\033[2J")

	clearScreen()

	// %-15s：表示字符串左对齐（- 是左对齐标志），占用 15 个字符宽度（不够则补空格）。
	// %-25s：表示字符串左对齐，占用 25 个字符宽度
	fmt.Printf("| %-15s | %-25s |\n", "City", "Time")
	fmt.Println("|-----------------|---------------------------|")

	for _, val := range citys {
		fmt.Printf("| %-15s | %-25s |\n", val.name, val.data)
	}

	time.Sleep(500 * time.Millisecond)
	// fmt.Print("\033[H\033[2J")
	// clearScreen()
}

func getTime(city *city) {
	conn, err := net.Dial("tcp", city.url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(city, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
