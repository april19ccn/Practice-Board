package main

import (
	"path/filepath"
	"strings"
)

// "a/b/c.go"

func basenameV1(str string) string {
	result := str
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '/' { // 比较的是字节值
			result = str[i+1:]
			break
		}
	}

	for i := len(result) - 1; i >= 0; i-- {
		if string(result[i]) == "." { // 字符串比较
			result = result[0:i]
			break
		}
	}

	return result
}

func basenameV2(str string) string {
	if strings.LastIndex(str, "/") >= 0 {
		str = str[strings.LastIndex(str, "/")+1:] // 获取最后一个/后面的字符串
	}
	if dotIndex := strings.LastIndex(str, "."); dotIndex >= 0 {
		str = str[:dotIndex] // 获取最后一个.前面的字符串
	}
	return str
}

func basenameV3(str string) string {
	_, file := filepath.Split(str) // 获取(路径)和文件名
	if dotIndex := strings.LastIndex(file, "."); dotIndex >= 0 {
		file = file[:dotIndex] // 获取最后一个.前面的字符串
	}
	return file
}

func basenameV4(str string) string {
	file := filepath.Base(str) // 获取(路径)和文件名
	if dotIndex := strings.LastIndex(file, "."); dotIndex >= 0 {
		file = file[:dotIndex] // 获取最后一个.前面的字符串
	}
	return file
}
