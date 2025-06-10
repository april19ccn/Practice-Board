// 不符合题意，没有使用flag
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func getSHA256(x []byte) [32]uint8 {
	return sha256.Sum256(x)
}

func getSHA384(x []byte) [48]uint8 {
	return sha512.Sum384(x)
}

func getSHA512(x []byte) [64]uint8 {
	return sha512.Sum512(x)
}

func main() {
	fmt.Println("输入要解析的字符串和模式「字符串 （256、384、512）」：")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Split(input.Text(), " ")
		if len(arr) == 1 {
			fmt.Printf("%x\n", getSHA256([]byte(arr[0])))
		} else {
			switch arr[1] {
			case "384":
				fmt.Printf("%x\n", getSHA384([]byte(arr[0])))
			case "512":
				fmt.Printf("%x\n", getSHA512([]byte(arr[0])))
			default:
				fmt.Printf("%x\n", getSHA256([]byte(arr[0])))
			}
		}
	}
}
