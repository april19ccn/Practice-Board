package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(184467440737095)
	}
}

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountFor(184467440737095)
	}
}
