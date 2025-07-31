package main

import (
	"fmt"
	"io"
	"os"
)

// func writeHeader(w io.Writer, contentType string) error {
// 	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
// 		return err
// 	}
// 	if _, err := w.Write([]byte(contentType)); err != nil {
// 		return err
// 	}
// 	// ...
// 	return nil
// }

// writeString writes s to w.
// If w has a WriteString method, it is invoked instead of w.Write.
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
	// ...
	return nil
}

// io.WriteString
func writeHeader1(w io.Writer, contentType string) error {
	if _, err := io.WriteString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := io.WriteString(w, contentType); err != nil {
		return err
	}
	// ...
	return nil
}

func main() {
	fmt.Fprint(os.Stdout, "Hello, world\n")
}

// 伪代码：
// package fmt

// func formatOneValue(x interface{}) string {
//     if err, ok := x.(error); ok {
//         return err.Error()
//     }
//     if str, ok := x.(Stringer); ok {
//         return str.String()
//     }
//     // ...all other types...
// }
