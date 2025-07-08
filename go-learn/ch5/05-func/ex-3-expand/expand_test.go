package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	got := expand("aabc foo, 123 foo fooo", strings.ToUpper)
	want := "aabc FOO, 123 FOO FOOo"
	if got != want {
		t.Errorf("expand() = %v, want %v", got, want)
	}
}

func TestExpandReplace(t *testing.T) {
	got := expandReplace("aabc foo, 123 foo fooo", strings.ToUpper)
	want := "aabc FOO, 123 FOO FOOo"
	if got != want {
		t.Errorf("expand() = %v, want %v", got, want)
	}
}
