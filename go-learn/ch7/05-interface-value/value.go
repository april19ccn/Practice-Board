package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	// var buf *bytes.Buffer
	// if debug {
	// 	buf = new(bytes.Buffer) // enable collection of output
	// }
	// f(buf) // NOTE: subtly incorrect!

	// 修复方案
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf)

	// fmt.Println(buf.(*bytes.Buffer).String()) // 当 debug = false 时，会引发 panic: interface conversion: io.Writer is nil, not *bytes.Buffer
	if debug {
		fmt.Println(buf.(*bytes.Buffer).String())
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil { // 当 debug = false 时，会引发 panic: runtime error: invalid memory address or nil pointer dereference
		out.Write([]byte("done!\n"))
		fmt.Println(out.(*bytes.Buffer).String())
	}
}
