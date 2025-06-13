package main

import "testing"

func BenchmarkTestMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMake()
	}
}
