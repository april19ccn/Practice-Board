// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 调用flag.Bool函数会创建一个新的对应布尔型标志参数的变量。
// 它有三个属性：
// 第一个是命令行标志参数的名字“n”
// 第二个是该标志参数的默认值（这里是false）
// 第三个是该标志参数对应的描述信息
var n = flag.Bool("n", false, "omit trailing newline")

// 调用flag.String函数将创建一个对应字符串类型的标志参数变量，同样包含命令行标志参数对应的参数名、默认值、和描述信息。
var sep = flag.String("s", " ", "separator")

func main() {
	// 解析 os.Args[1:] 中的命令行标志。必须在定义了所有标志之后、程序访问标志之前调用。
	flag.Parse()

	// 程序中的sep和n变量分别是指向对应命令行标志参数变量的指针，因此必须用*sep和*n形式的指针语法间接引用它们。
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
		fmt.Println("zzz")
	}
}

// 测试方法
// go run .\echo.go -s / a bc def
// output：
// a/bc/def
// zzz

// go run .\echo.go -n a bc def
// output：
// a/bc/def

// go run .\echo.go -help
// output：
// Usage of C:\Users\Administrator\AppData\Local\Temp\go-build3520654578\b001\exe\echo.exe:
//   -n    omit trailing newline
//   -s string
//         separator (default " ")
