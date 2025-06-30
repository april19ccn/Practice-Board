// 练习 4.13： 使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。
// 编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type OMDBParams struct {
	Title string `json:"t"`
}

type OMDBResult struct {
	Title  string
	Poster string
}

var MovieName = flag.String("n", "", "Movie name")

const OMDBAPI = "http://www.omdbapi.com/?apikey="

// GetOMDBDir 返回OMDB下载目录
func GetOMDBDir() (string, error) {
	dir, err := os.Getwd() // 获取当前工作目录
	fmt.Println(dir)
	if err != nil {
		return "", err
	}

	// targetDir := filepath.Join(dir, "ch4", "05-json", "ex-movies", "poster")
	targetDir := filepath.Join(dir, "poster")
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", err
	}

	return targetDir, nil
}

func GetOMDB(title string) (*OMDBResult, error) {
	t := url.QueryEscape(title)
	req, err := http.NewRequest("GET", OMDBAPI+"&t="+t, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result OMDBResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	fmt.Println(result)
	return &result, err
}

func GetPoster(OMDB *OMDBResult) error {
	posterUrl := OMDB.Poster

	pots := strings.Split(posterUrl, ".")
	fileName := OMDB.Title + "." + pots[len(pots)-1]

	path, err := GetOMDBDir()
	if err != nil {
		log.Fatal(err)
	}
	filePath := filepath.Join(path, fileName)
	fmt.Println(filePath)

	// 获取封面
	resp, err := http.Get(posterUrl)
	if err != nil {
		return fmt.Errorf("无法获取文件: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("无法创建保存文件: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("无法写入文件: %v", err)
	}

	return nil
}

func main() {
	flag.Parse()

	// _, err := GetOMDB("Blade Runner")
	res, err := GetOMDB(*MovieName)
	if err != nil {
		log.Fatal(err)
	}

	err = GetPoster(res)
	if err != nil {
		log.Fatal(err)
	}
}

// go run .\movies.go -n "Laputa: Castle in the Sky"
// Laputa: Castle in the Sky (1986).jpg 文件名非法
