// 练习 10.1： 扩展jpeg程序，以支持任意图像格式之间的相互转换，
// 使用image.Decode检测支持的格式类型，然后通过flag命令行标志参数选择输出的格式。

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"strings"

	"image/gif"
	"image/png" // register PNG decoder

	"golang.org/x/image/tiff"

	"io"
	"os"
)

var inputFile = flag.String("i", "", "Input File")
var targetFormat = flag.String("t", "png", "JPEG、GIF、PNG、TIFF")

func main() {
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("请导入文件！")
		return
	}

	file, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := toTargetFormat(bytes.NewReader(file), os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func toTargetFormat(in io.Reader, out io.Writer) error {
	// 解析图片格式
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	// 转换格式
	switch strings.ToLower(*targetFormat) {
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	case "png":
		return png.Encode(out, img)
	case "tiff":
		return tiff.Encode(out, img, &tiff.Options{})
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	default:
		return fmt.Errorf("不支持目标格式！")
	}
}
