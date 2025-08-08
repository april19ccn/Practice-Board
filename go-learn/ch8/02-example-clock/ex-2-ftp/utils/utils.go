package utils

import (
	"log"
	"os"
	"os/user"
)

// 获取工作目录，currentPath 为当前工作目录，homeDir 为home目录
func GetWorkPath() (currentPath string, homeDir string) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir = u.HomeDir

	return currentPath, homeDir
}
