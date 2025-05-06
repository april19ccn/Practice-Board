package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCommandName(t *testing.T) {
	t.Run("Test Command Name", func(t *testing.T) {
		buffer := bytes.Buffer{}
		printCommandName(&buffer)

		got := buffer.String()
		want := "ex-echo.test.exe"

		if !strings.Contains(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkConnectStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		connectStringJoin([]string{"test1", "test2", "test3", "test4", "test5", "test6"})
	}
}

func BenchmarkConnectStringFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		connectStringFor([]string{"test1", "test2", "test3", "test4", "test5", "test6"})
	}
}
