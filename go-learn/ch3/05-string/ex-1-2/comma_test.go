package main

import "testing"

func TestComma(t *testing.T) {
	got := comma("45689321689")
	want := "45,689,321,689"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestComma2(t *testing.T) {
	got := comma("45689321689.88888")
	want := "45,689,321,689.88888"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestComma3(t *testing.T) {
	got := comma("+45689321689.88888")
	want := "+45,689,321,689.88888"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestComma4(t *testing.T) {
	got := comma("+845689321689.88888")
	want := "+845,689,321,689.88888"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma("45689321689.88888")
	}
}
