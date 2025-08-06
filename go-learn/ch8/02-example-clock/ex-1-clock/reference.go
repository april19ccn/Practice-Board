package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var citySlice []city

type city struct {
	name string
	data *data
}
type data string

func main() {

	var wg sync.WaitGroup
	wg.Add(len(os.Args[1:]))

	for i, val := range os.Args[1:] {
		line := strings.Split(val, "=")
		name, address := line[0], line[1]
		citySlice = append(citySlice, city{name: name, data: new(data)})

		go getTime(address, citySlice[i].data, &wg)

	}
	for {
		// 使用ANSI 转义序列操作终端的指令
		// \033[H：将光标移动到终端的左上角（"Home" 位置）
		// \033[2J：清除终端屏幕上的所有内容
		fmt.Print("\033[H\033[2J")

		fmt.Printf("| %-15s | %-25s |\n", "City", "Time")
		fmt.Println("|-----------------|---------------------------|")

		for _, val := range citySlice {
			printRow(val.name, val.data)
		}

		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J")
	}
	wg.Wait()

}

func printRow(city string, time *data) {
	fmt.Printf("| %-15s | %-25s |\n", city, string(*time))
}

func (d *data) String() string {
	return string(*d)
}

func (d *data) Write(p []byte) (n int, err error) {
	*d = data(p)
	return len(p), nil
}

func getTime(address string, data *data, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(data, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}
}
