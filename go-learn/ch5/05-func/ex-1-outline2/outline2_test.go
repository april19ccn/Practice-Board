package main

import (
	"bytes"
	"os"
	"testing"
)

func TestOutline(t *testing.T) {
	var got bytes.Buffer
	err := outline(&got, true, "./template.html")
	if err != nil {
		t.Fatal(err)
	}

	want, err := os.ReadFile("./template-res.html")
	if err != nil {
		t.Fatal(err)
	}

	if got.String() != string(want) {
		t.Errorf("outline() = %v, want %v", got.String(), string(want))
	}
}
