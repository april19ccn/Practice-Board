// intsToString is like fmt.Sprint(values) but adds commas.
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func intsToStringV1(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func intsToStringV2(values []int) string {
	buf := []byte("[")
	for i, v := range values {
		if i > 0 {
			buf = append(buf, []byte(", ")...)
		}
		buf = append(buf, []byte(strconv.Itoa(v))...)
	}
	buf = append(buf, []byte("]")...)
	return string(buf)
}

func intsToString(values []int) string {
	buf := make([]byte, 0, 2+(len(values)-1)*2)
	buf = append(buf, '[')
	for i, v := range values {
		if i > 0 {
			buf = append(buf, ',', ' ')
		}
		buf = append(buf, []byte(strconv.Itoa(v))...)
	}
	buf = append(buf, ']')
	return string(buf)
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
