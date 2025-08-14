package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestOutGoSVG(t *testing.T) {
	OutGoSVG()

	got, err := os.Open("./test.svg")
	if err != nil {
		t.Fatal(err)
	}
	defer got.Close()

	want, err := os.Open("./out.svg")
	if err != nil {
		t.Fatal(err)
	}
	defer want.Close()

	gotContent, _ := io.ReadAll(got)
	wantContent, _ := io.ReadAll(want)
	if !bytes.Equal(gotContent, wantContent) {
		t.Errorf("SVG content does not match expected output")
	}
}

func TestOutGoSVG2(t *testing.T) {
	OutGoSVG2()

	got, err := os.Open("./test.svg")
	if err != nil {
		t.Fatal(err)
	}
	defer got.Close()

	want, err := os.Open("./out.svg")
	if err != nil {
		t.Fatal(err)
	}
	defer want.Close()

	gotContent, _ := io.ReadAll(got)
	wantContent, _ := io.ReadAll(want)
	if !bytes.Equal(gotContent, wantContent) {
		t.Errorf("SVG content does not match expected output")
	}
}

// goos: windows
// goarch: amd64
// pkg: example/learn/ch8/05-thumbnail/ex-02-surface
// cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
// BenchmarkOutSVG-6   	      13	  88581546 ns/op	  640861 B/op	   80003 allocs/op
// PASS
// ok  	example/learn/ch8/05-thumbnail/ex-02-surface	1.602s
func BenchmarkOutSVG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OutSVG()
	}
}

func BenchmarkOutGoSVG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OutGoSVG()
	}
}

func BenchmarkOutGoSVG2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OutGoSVG2()
	}
}
