package zip

import (
	"archive/zip"
	"example/learn/ch10/05-anonymous-import/ex-2-zip/core"
	"fmt"
	"io"
	"log"
)

func Decode(path string) {
	zipFile, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	// 遍历 ZIP 文件中的每个文件
	for _, file := range zipFile.File {
		fmt.Printf("文件名: %s\n", file.Name)
		fmt.Printf("文件大小: %d bytes\n", file.UncompressedSize64)
		fmt.Printf("修改时间: %s\n", file.Modified.Format("2006-01-02 15:04:05"))
		fmt.Printf("是否是目录: %v\n", file.Mode().IsDir())

		// 如果不是目录，则读取文件内容
		if !file.Mode().IsDir() {
			// 打开 ZIP 文件中的文件
			rc, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer rc.Close()

			// 读取文件内容
			content, err := io.ReadAll(rc)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("文件内容:\n%s\n", string(content))
		}
		fmt.Println("----------------------------------")
	}
}

func init() {
	core.Register("zip", Decode)
}
