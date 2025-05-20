package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(184467440737095)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(184467440737095)
	}
}

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountFor(184467440737095)
	}
}

func BenchmarkPopCountDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountDel(184467440737095)
	}
}
