package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestDiffCount(t *testing.T) {
	dataList := []struct {
		id int8
		x  []byte
		y  []byte
	}{
		{
			id: 1,
			x:  []byte("x"),
			y:  []byte("X"),
		},
		{
			id: 1,
			x:  []byte("x123h12"),
			y:  []byte("X456"),
		},
	}

	for _, v := range dataList {
		got1 := diffCount(sha256.Sum256(v.x), sha256.Sum256(v.y))
		got2 := popCount(sha256.Sum256(v.x), sha256.Sum256(v.y))

		fmt.Println(v.id, got1, got2)

		if got1 != got2 {
			t.Errorf("id: %d, got: %d, want: %d", v.id, got1, got2)
		}
	}
}
