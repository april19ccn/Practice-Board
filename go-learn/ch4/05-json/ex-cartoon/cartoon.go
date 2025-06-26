// 练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。
// 例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。
// 下载每个链接（只下载一次）然后创建一个离线索引。
// 编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type XKCD struct {
	Num        int    `json:"num"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Alt        string `json:"alt"`        // 替代文本 对图片、漫画等视觉内容的文字说明，用于无障碍访问（如屏幕阅读器读取）
	Transcript string `json:"transcript"` // 文本转录 详细记录内容中的对话、场景描述或情节流程（如漫画分镜脚本）
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	News       string `json:"news"`       // 新闻内容 存储与内容相关的新闻文本或描述性信息
	SafeTitle  string `json:"safe_title"` // 安全标题 提供一个简洁、无特殊字符的标题，用于系统识别或文件命名
	Link       string `json:"link"`       // 存储该内容的外部链接（如网页地址），便于用户跳转访问。
}

var XKCDNum = flag.Int("n", 0, "XKCD number")

const XKCDURL = "https://xkcd.com/%d/info.0.json"

func GetXKCDDir() (string, error) {
	dir, err := os.Getwd()
	fmt.Println(dir)
	if err != nil {
		return "", err
	}

	// targetDir := filepath.Join(dir, "ch4", "05-json", "ex-cartoon", "xkcd")
	targetDir := filepath.Join(dir, "xkcd")
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", err
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", err
	}

	return targetDir, nil
}

func GetOfflineXKCD(num int) (*XKCD, error) {
	targetDir, err := GetXKCDDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(targetDir, fmt.Sprintf("%d.json", num))
	_, err = os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, err // 文件不存在
		}
		return nil, fmt.Errorf("检查文件时出错: %w", err) // 其他错误
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var xkcd XKCD
	if err := json.Unmarshal(data, &xkcd); err != nil {
		return nil, err
	}

	return &xkcd, nil
}

func GetOnlineXKCD(num int) (*XKCD, error) {
	var xkcd XKCD
	resp, err := http.Get(fmt.Sprintf(XKCDURL, num))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&xkcd); err != nil {
		return nil, err
	}
	return &xkcd, nil
}

func CreateXKCDFile(num int, XKCD *XKCD) error {
	targetDir, err := GetXKCDDir()
	if err != nil {
		return err
	}

	filePath := filepath.Join(targetDir, fmt.Sprintf("%d.json", num))
	data, err := os.Create(filePath)
	if err != nil {
		return err
	}

	formattedJSON, err := json.MarshalIndent(XKCD, "", "  ")
	if err != nil {
		fmt.Println("JSON 格式化失败:", err)
		return err
	}

	_, err = data.WriteString(string(formattedJSON))
	if err != nil {
		return err
	}

	return nil
}

func GetXKCD(num int) {
	data, err := GetOfflineXKCD(num)
	if err != nil {
		fmt.Println(err)
	}

	if data == nil {
		fmt.Println("获取在线数据")
		data, err = GetOnlineXKCD(num)
		if err != nil {
			fmt.Println(err)
		}
		CreateXKCDFile(num, data)
	}

	fmt.Println(data)
}

func main() {
	flag.Parse()

	GetXKCD(*XKCDNum)
}
