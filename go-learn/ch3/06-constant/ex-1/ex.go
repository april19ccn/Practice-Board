package main

import (
	"fmt"
	"math/big"
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)

	// ZB 和 YB 的值超出了 Go 语言中 int 类型的范围（在 64 位系统中，int 最大约 9e18）
	// 使用 big.Int 计算并打印 ZB 和 YB
	kb := big.NewInt(1000)
	zb := new(big.Int).Exp(kb, big.NewInt(7), nil) // 1000^7
	yb := new(big.Int).Exp(kb, big.NewInt(8), nil) // 1000^8
	fmt.Println(zb)
	fmt.Println(yb)
}
