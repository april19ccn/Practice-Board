package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var SHAType = flag.Int("type", 256, "请选择SHA解析方式,256/384/512")

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
	flag.Parse()

	fmt.Println(flag.Args()[0])
	fmt.Println(os.Args[1])
	fmt.Println(*SHAType)

	switch *SHAType {
	case 384:
		fmt.Printf("%x\n", getSHA384([]byte(flag.Args()[0])))
	case 512:
		fmt.Printf("%x\n", getSHA512([]byte(flag.Args()[0])))
	default:
		fmt.Printf("%x\n", getSHA256([]byte(flag.Args()[0])))
	}
}

// go run .\printSHA.go -type 512 x (不能写成 x -type 512)
