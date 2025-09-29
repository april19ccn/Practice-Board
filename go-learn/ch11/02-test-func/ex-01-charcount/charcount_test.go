package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestCharcount(t *testing.T) {
	testList := []struct {
		id   int
		data string
		want string
	}{
		{
			id:   1,
			data: "hui\nyu\n123\n",
			want: "\nlen\tcount\n1\t11\n2\t0\n3\t0\n4\t0\n",
		},
	}

	for _, v := range testList {
		fmt.Println("------" + strconv.Itoa(v.id) + "------")
		var buf bytes.Buffer
		Charcount(strings.NewReader(v.data), &buf)
		fmt.Println(buf.String())
		if v.want != buf.String() {
			t.Errorf("Fail id = %d", v.id)
		}
	}
}
