package main

import "testing"

func TestComma(t *testing.T) {
	got := comma("45689321689")
	want := "45,689,321,689"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
