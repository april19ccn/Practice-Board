package Tar

import (
	"archive/tar"
	"example/learn/ch10/05-anonymous-import/ex-2-zip/core"
	"fmt"
	"io"
	"log"
	"os"
)

func Decode(path string) {
	// 打开 tar 文件
	tarFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer tarFile.Close()

	// 创建 tar reader
	tarReader := tar.NewReader(tarFile)

	// 遍历 tar 文件中的每个文件
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// 文件结束
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("文件名: %s\n", header.Name)
		fmt.Printf("文件大小: %d bytes\n", header.Size)
		fmt.Printf("修改时间: %s\n", header.ModTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("文件模式: %v\n", header.Mode)
		fmt.Printf("是否是目录: %v\n", header.Typeflag == tar.TypeDir)

		// 如果不是目录，则读取文件内容
		if header.Typeflag != tar.TypeDir {
			// 读取文件内容
			content, err := io.ReadAll(tarReader)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("文件内容:\n%s\n", string(content))
		}
		fmt.Println("----------------------------------")
	}
}

func init() {
	core.Register("tar", Decode)
}
