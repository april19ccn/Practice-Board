// 练习 10.2： 设计一个通用的压缩文件读取框架，用来读取ZIP（archive/zip）和POSIX tar（archive/tar）格式压缩的文档。
// 使用类似上面的注册技术来扩展支持不同的压缩格式，然后根据需要通过匿名导入选择导入要支持的压缩格式的驱动包。
package main

import (
	"example/learn/ch10/05-anonymous-import/ex-2-zip/core"
	_ "example/learn/ch10/05-anonymous-import/ex-2-zip/parseTar"
	_ "example/learn/ch10/05-anonymous-import/ex-2-zip/parseZip"
	"fmt"
	"path/filepath"
)

func main() {
	// core.Decode("zip", `D:\Star_Code\A_Break\-Test\practice\go-learn\ch10\05-anonymous-import\ex-2-zip\test2.zip`)
	absolutePath, err := filepath.Abs(`ch10/05-anonymous-import/ex-2-zip/test1.tar`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("绝对路径:", absolutePath)
	core.Decode("tar", absolutePath)
}
