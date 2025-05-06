package main

import "testing"

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
